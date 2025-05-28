package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"io"
)

// Structs representing varying types of LineStrings
type LineString struct {
	Points []Point
}

type LineStringZ struct {
	Points []PointZ
}

type LineStringM struct {
	Points []PointM
}

type LineStringZM struct {
	Points []PointZM
}

type LineStringS struct {
	SRID   int32
	Points []Point
}

type LineStringZS struct {
	SRID   int32
	Points []PointZ
}

type LineStringMS struct {
	SRID   int32
	Points []PointM
}

type LineStringZMS struct {
	SRID   int32
	Points []PointZM
}

// Implement SRIDGeometry interface for SRID types
func (ls *LineStringS) GetSRID() int32     { return ls.SRID }
func (ls *LineStringS) SetSRID(srid int32) { ls.SRID = srid }

func (ls *LineStringZS) GetSRID() int32     { return ls.SRID }
func (ls *LineStringZS) SetSRID(srid int32) { ls.SRID = srid }

func (ls *LineStringMS) GetSRID() int32     { return ls.SRID }
func (ls *LineStringMS) SetSRID(srid int32) { ls.SRID = srid }

func (ls *LineStringZMS) GetSRID() int32     { return ls.SRID }
func (ls *LineStringZMS) SetSRID(srid int32) { ls.SRID = srid }

// Implement CollectionGeometry interface for all LineString types
func (ls *LineString) GetElementCount() uint32    { return getElementCountHelper(ls.Points) }
func (ls *LineStringZ) GetElementCount() uint32   { return getElementCountHelper(ls.Points) }
func (ls *LineStringM) GetElementCount() uint32   { return getElementCountHelper(ls.Points) }
func (ls *LineStringZM) GetElementCount() uint32  { return getElementCountHelper(ls.Points) }
func (ls *LineStringS) GetElementCount() uint32   { return getElementCountHelper(ls.Points) }
func (ls *LineStringZS) GetElementCount() uint32  { return getElementCountHelper(ls.Points) }
func (ls *LineStringMS) GetElementCount() uint32  { return getElementCountHelper(ls.Points) }
func (ls *LineStringZMS) GetElementCount() uint32 { return getElementCountHelper(ls.Points) }

func (ls *LineString) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringZ) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringM) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringZM) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringS) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringZS) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringMS) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineStringZMS) WriteElements(buffer *bytes.Buffer) error {
	return writeElementsHelper(ls.Points, buffer)
}

func (ls *LineString) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[Point](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringZ) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[PointZ](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringM) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[PointM](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringZM) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[PointZM](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringS) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[Point](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringZS) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[PointZ](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringMS) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[PointM](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

func (ls *LineStringZMS) ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error {
	points, err := readElementsHelper[PointZM](reader, byteOrder, count)
	if err != nil {
		return err
	}
	ls.Points = points
	return nil
}

/** LineString functions **/
func (ls *LineString) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineString) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineString) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineString) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXY, false)
}

/** LineStringZ functions **/
func (ls *LineStringZ) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringZ) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringZ) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringZ) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXYZ, false)
}

/** LineStringM functions **/
func (ls *LineStringM) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringM) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringM) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringM) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXYM, false)
}

/** LineStringZM functions **/
func (ls *LineStringZM) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringZM) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringZM) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringZM) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXYZM, false)
}

/** LineStringS functions **/
func (ls *LineStringS) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringS) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringS) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringS) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXY, true)
}

/** LineStringZS functions **/
func (ls *LineStringZS) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringZS) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringZS) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringZS) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXYZ, true)
}

/** LineStringMS functions **/
func (ls *LineStringMS) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringMS) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringMS) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringMS) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXYM, true)
}

/** LineStringZMS functions **/
func (ls *LineStringZMS) Scan(value interface{}) error {
	return scanGeometryHelper(ls, value)
}

func (ls LineStringZMS) Value() (driver.Value, error) {
	return valueGeometryHelper(&ls)
}

func (ls LineStringZMS) Write(buffer *bytes.Buffer) error {
	return WriteGeometryCollection(buffer, ls.GetElementCount(), func(buf *bytes.Buffer) error {
		return ls.WriteElements(buf)
	})
}

func (ls LineStringZMS) GetType() uint32 {
	return BuildWKBType(WKBLineString, CoordXYZM, true)
}
