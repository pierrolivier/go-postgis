package postgis

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"io"
)

// Core geometry interface that all PostGIS types must implement
type Geometry interface {
	sql.Scanner
	driver.Valuer

	GetType() uint32
	Write(*bytes.Buffer) error
}

// SRIDGeometry interface for geometries that have SRID
type SRIDGeometry interface {
	GetSRID() int32
	SetSRID(srid int32)
}

// PointReader interface for types that can read point data
type PointReader interface {
	ReadPoint(reader io.Reader, byteOrder binary.ByteOrder) error
}

// PointWriter interface for types that can write point data
type PointWriter interface {
	WritePoint(buffer *bytes.Buffer) error
}

// CollectionGeometry interface for geometries that contain multiple elements
type CollectionGeometry interface {
	GetElementCount() uint32
	WriteElements(buffer *bytes.Buffer) error
	ReadElements(reader io.Reader, byteOrder binary.ByteOrder, count uint32) error
}

// Helper interface for common point operations (internal use)
type PointGeometry interface {
	Geometry
	scanPoint(value interface{}) error
	valuePoint() (driver.Value, error)
}
