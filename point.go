package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"io"
)

// Structs representing varying types of points
type Point struct {
	X, Y float64
}

type PointZ struct {
	X, Y, Z float64
}

type PointM struct {
	X, Y, M float64
}

type PointZM struct {
	X, Y, Z, M float64
}

type PointS struct {
	SRID int32
	X, Y float64
}

type PointZS struct {
	SRID    int32
	X, Y, Z float64
}

type PointMS struct {
	SRID    int32
	X, Y, M float64
}

type PointZMS struct {
	SRID       int32
	X, Y, Z, M float64
}

// Implement SRIDGeometry interface for SRID types
func (p *PointS) GetSRID() int32     { return p.SRID }
func (p *PointS) SetSRID(srid int32) { p.SRID = srid }

func (p *PointZS) GetSRID() int32     { return p.SRID }
func (p *PointZS) SetSRID(srid int32) { p.SRID = srid }

func (p *PointMS) GetSRID() int32     { return p.SRID }
func (p *PointMS) SetSRID(srid int32) { p.SRID = srid }

func (p *PointZMS) GetSRID() int32     { return p.SRID }
func (p *PointZMS) SetSRID(srid int32) { p.SRID = srid }

// Implement PointReader interface for SRID types (they need special handling)
func (p *PointS) ReadPoint(reader io.Reader, byteOrder binary.ByteOrder) error {
	if err := binary.Read(reader, byteOrder, &p.X); err != nil {
		return err
	}
	return binary.Read(reader, byteOrder, &p.Y)
}

func (p *PointZS) ReadPoint(reader io.Reader, byteOrder binary.ByteOrder) error {
	if err := binary.Read(reader, byteOrder, &p.X); err != nil {
		return err
	}
	if err := binary.Read(reader, byteOrder, &p.Y); err != nil {
		return err
	}
	return binary.Read(reader, byteOrder, &p.Z)
}

func (p *PointMS) ReadPoint(reader io.Reader, byteOrder binary.ByteOrder) error {
	if err := binary.Read(reader, byteOrder, &p.X); err != nil {
		return err
	}
	if err := binary.Read(reader, byteOrder, &p.Y); err != nil {
		return err
	}
	return binary.Read(reader, byteOrder, &p.M)
}

func (p *PointZMS) ReadPoint(reader io.Reader, byteOrder binary.ByteOrder) error {
	if err := binary.Read(reader, byteOrder, &p.X); err != nil {
		return err
	}
	if err := binary.Read(reader, byteOrder, &p.Y); err != nil {
		return err
	}
	if err := binary.Read(reader, byteOrder, &p.Z); err != nil {
		return err
	}
	return binary.Read(reader, byteOrder, &p.M)
}

/** Point functions **/
func (p *Point) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p Point) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p Point) Write(buffer *bytes.Buffer) error {
	return binary.Write(buffer, binary.LittleEndian, &p)
}

func (p Point) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXY, false)
}

/** PointZ functions **/
func (p *PointZ) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointZ) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointZ) Write(buffer *bytes.Buffer) error {
	return binary.Write(buffer, binary.LittleEndian, &p)
}

func (p PointZ) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXYZ, false)
}

/** PointM functions **/
func (p *PointM) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointM) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointM) Write(buffer *bytes.Buffer) error {
	return binary.Write(buffer, binary.LittleEndian, &p)
}

func (p PointM) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXYM, false)
}

/** PointZM functions **/
func (p *PointZM) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointZM) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointZM) Write(buffer *bytes.Buffer) error {
	return binary.Write(buffer, binary.LittleEndian, &p)
}

func (p PointZM) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXYZM, false)
}

/** PointS functions **/
func (p *PointS) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointS) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointS) Write(buffer *bytes.Buffer) error {
	// For SRID types, don't write SRID here - it's handled by WriteEWKB
	return binary.Write(buffer, binary.LittleEndian, &struct{ X, Y float64 }{p.X, p.Y})
}

func (p PointS) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXY, true)
}

/** PointZS functions **/
func (p *PointZS) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointZS) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointZS) Write(buffer *bytes.Buffer) error {
	// For SRID types, don't write SRID here - it's handled by WriteEWKB
	return binary.Write(buffer, binary.LittleEndian, &struct{ X, Y, Z float64 }{p.X, p.Y, p.Z})
}

func (p PointZS) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXYZ, true)
}

/** PointMS functions **/
func (p *PointMS) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointMS) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointMS) Write(buffer *bytes.Buffer) error {
	// For SRID types, don't write SRID here - it's handled by WriteEWKB
	return binary.Write(buffer, binary.LittleEndian, &struct{ X, Y, M float64 }{p.X, p.Y, p.M})
}

func (p PointMS) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXYM, true)
}

/** PointZMS functions **/
func (p *PointZMS) Scan(value interface{}) error {
	return scanGeometryHelper(p, value)
}

func (p PointZMS) Value() (driver.Value, error) {
	return valueGeometryHelper(&p)
}

func (p PointZMS) Write(buffer *bytes.Buffer) error {
	// For SRID types, don't write SRID here - it's handled by WriteEWKB
	return binary.Write(buffer, binary.LittleEndian, &struct{ X, Y, Z, M float64 }{p.X, p.Y, p.Z, p.M})
}

func (p PointZMS) GetType() uint32 {
	return BuildWKBType(WKBPoint, CoordXYZM, true)
}
