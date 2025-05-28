package postgis

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// Byte order constants for WKB
const (
	wkbXDR byte = 0
	wkbNDR byte = 1
)

// GeometryType constants for WKB geometry types
const (
	WKBPoint           uint32 = 1
	WKBLineString      uint32 = 2
	WKBPolygon         uint32 = 3
	WKBMultiPoint      uint32 = 4
	WKBMultiLineString uint32 = 5
	WKBMultiPolygon    uint32 = 6

	// Flags for coordinate dimensions
	WKBZFlag    uint32 = 0x80000000
	WKBMFlag    uint32 = 0x40000000
	WKBSRIDFlag uint32 = 0x20000000
)

// CoordinateType represents the type of coordinates (2D, Z, M, ZM)
type CoordinateType int

const (
	CoordXY CoordinateType = iota
	CoordXYZ
	CoordXYM
	CoordXYZM
)

// GeometryInfo holds metadata about a geometry type
type GeometryInfo struct {
	BaseType  uint32
	CoordType CoordinateType
	HasSRID   bool
}

// GetGeometryInfo extracts geometry information from a WKB type code
func GetGeometryInfo(wkbType uint32) GeometryInfo {
	info := GeometryInfo{
		BaseType: wkbType & 0x1FFFFFFF, // Remove flags to get base type
		HasSRID:  (wkbType & WKBSRIDFlag) != 0,
	}

	// Determine coordinate type from flags
	hasZ := (wkbType & WKBZFlag) != 0
	hasM := (wkbType & WKBMFlag) != 0

	switch {
	case hasZ && hasM:
		info.CoordType = CoordXYZM
	case hasZ:
		info.CoordType = CoordXYZ
	case hasM:
		info.CoordType = CoordXYM
	default:
		info.CoordType = CoordXY
	}

	return info
}

// BuildWKBType constructs a WKB type code from geometry information
func BuildWKBType(baseType uint32, coordType CoordinateType, hasSRID bool) uint32 {
	wkbType := baseType

	switch coordType {
	case CoordXYZ:
		wkbType |= WKBZFlag
	case CoordXYM:
		wkbType |= WKBMFlag
	case CoordXYZM:
		wkbType |= WKBZFlag | WKBMFlag
	}

	if hasSRID {
		wkbType |= WKBSRIDFlag
	}

	return wkbType
}

// Since Postgres by default returns hex encoded strings we need to first get bytes back
func DecodeEWKB(value interface{}) (io.Reader, error) {
	var ewkb []byte
	var err error

	switch v := value.(type) {
	case string:
		// For pgx, decode the hex-encoded string into bytes
		ewkb, err = hex.DecodeString(v)
		if err != nil {
			return nil, err
		}
	case []byte:
		// For lib/pq, cast it to string and decode the hex-encoded string into bytes
		ewkb, err = hex.DecodeString(string(v))
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported type: %T", value)
	}

	return bytes.NewReader(ewkb), nil
}

// EncodeEWKB encodes a buffer to hex string
func EncodeEWKB(buffer *bytes.Buffer) string {
	return hex.EncodeToString(buffer.Bytes())
}

// WriteEWKB writes a geometry to EWKB format
func WriteEWKB(g Geometry) (*bytes.Buffer, error) {
	buffer := bytes.NewBuffer(nil)

	// Set our endianness
	if err := binary.Write(buffer, binary.LittleEndian, wkbNDR); err != nil {
		return nil, err
	}

	// Write geometry type
	if err := binary.Write(buffer, binary.LittleEndian, g.GetType()); err != nil {
		return nil, err
	}

	// Write SRID if present
	if sridGeom, ok := g.(SRIDGeometry); ok {
		if err := binary.Write(buffer, binary.LittleEndian, sridGeom.GetSRID()); err != nil {
			return nil, err
		}
	}

	// Write geometry data
	if err := g.Write(buffer); err != nil {
		return nil, err
	}

	return buffer, nil
}

// ReadEWKB reads a geometry from EWKB format
func ReadEWKB(reader io.Reader, g Geometry) error {
	var byteOrder binary.ByteOrder
	var wkbByteOrder byte
	var wkbType uint32

	// Read byte order
	if err := binary.Read(reader, binary.LittleEndian, &wkbByteOrder); err != nil {
		return err
	}

	// Decide byte order
	switch wkbByteOrder {
	case wkbXDR:
		byteOrder = binary.BigEndian
	case wkbNDR:
		byteOrder = binary.LittleEndian
	default:
		return errors.New("unsupported byte order")
	}

	// Read geometry type
	if err := binary.Read(reader, byteOrder, &wkbType); err != nil {
		return err
	}

	info := GetGeometryInfo(wkbType)

	// Read SRID if present
	if info.HasSRID {
		if sridGeom, ok := g.(SRIDGeometry); ok {
			var srid int32
			if err := binary.Read(reader, byteOrder, &srid); err != nil {
				return err
			}
			sridGeom.SetSRID(srid)
		} else {
			return fmt.Errorf("geometry type %T does not support SRID but EWKB contains SRID", g)
		}
	}

	// Read geometry data using specialized readers
	return ReadGeometryData(reader, byteOrder, g, info)
}

// ReadGeometryData reads geometry-specific data
func ReadGeometryData(reader io.Reader, byteOrder binary.ByteOrder, g Geometry, info GeometryInfo) error {
	switch info.BaseType {
	case WKBPoint:
		// Points are simple - just read the coordinate data
		if pointReader, ok := g.(PointReader); ok {
			return pointReader.ReadPoint(reader, byteOrder)
		}
		// Fallback to binary.Read for simple point types
		return binary.Read(reader, byteOrder, g)

	case WKBLineString:
		// LineStrings need to read point count first, then points
		if collGeom, ok := g.(CollectionGeometry); ok {
			var count uint32
			if err := binary.Read(reader, byteOrder, &count); err != nil {
				return err
			}
			return collGeom.ReadElements(reader, byteOrder, count)
		}
		return fmt.Errorf("geometry type %T does not implement CollectionGeometry interface", g)

	default:
		return fmt.Errorf("unsupported geometry type: %d", info.BaseType)
	}
}

// WriteGeometryCollection writes a collection of points (for LineString, etc.)
func WriteGeometryCollection(buffer *bytes.Buffer, count uint32, writeElements func(*bytes.Buffer) error) error {
	// Write count
	if err := binary.Write(buffer, binary.LittleEndian, count); err != nil {
		return err
	}

	// Write elements
	return writeElements(buffer)
}

// ReadPointCollection reads a collection of points with the specified coordinate type
func ReadPointCollection(reader io.Reader, byteOrder binary.ByteOrder, count uint32, coordType CoordinateType) (interface{}, error) {
	switch coordType {
	case CoordXY:
		points := make([]Point, count)
		for i := uint32(0); i < count; i++ {
			if err := binary.Read(reader, byteOrder, &points[i]); err != nil {
				return nil, err
			}
		}
		return points, nil

	case CoordXYZ:
		points := make([]PointZ, count)
		for i := uint32(0); i < count; i++ {
			if err := binary.Read(reader, byteOrder, &points[i]); err != nil {
				return nil, err
			}
		}
		return points, nil

	case CoordXYM:
		points := make([]PointM, count)
		for i := uint32(0); i < count; i++ {
			if err := binary.Read(reader, byteOrder, &points[i]); err != nil {
				return nil, err
			}
		}
		return points, nil

	case CoordXYZM:
		points := make([]PointZM, count)
		for i := uint32(0); i < count; i++ {
			if err := binary.Read(reader, byteOrder, &points[i]); err != nil {
				return nil, err
			}
		}
		return points, nil

	default:
		return nil, fmt.Errorf("unsupported coordinate type: %d", coordType)
	}
}
