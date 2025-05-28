package postgis

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"io"
)

// Generic helper functions to reduce code duplication across geometry types

// scanGeometryHelper provides common Scan implementation for all geometry types
func scanGeometryHelper(g Geometry, value interface{}) error {
	reader, err := DecodeEWKB(value)
	if err != nil {
		return err
	}
	return ReadEWKB(reader, g)
}

// valueGeometryHelper provides common Value implementation for all geometry types
func valueGeometryHelper(g Geometry) (driver.Value, error) {
	buffer, err := WriteEWKB(g)
	if err != nil {
		return nil, err
	}
	return EncodeEWKB(buffer), nil
}

// writeElementsHelper provides common WriteElements implementation for collection types
func writeElementsHelper[T interface{ Write(*bytes.Buffer) error }](points []T, buffer *bytes.Buffer) error {
	for _, point := range points {
		if err := point.Write(buffer); err != nil {
			return err
		}
	}
	return nil
}

// readElementsHelper provides common ReadElements implementation for collection types
func readElementsHelper[T any](reader io.Reader, byteOrder binary.ByteOrder, count uint32) ([]T, error) {
	points := make([]T, count)
	for i := uint32(0); i < count; i++ {
		if err := binary.Read(reader, byteOrder, &points[i]); err != nil {
			return nil, err
		}
	}
	return points, nil
}

// getElementCountHelper provides common GetElementCount implementation
func getElementCountHelper[T any](points []T) uint32 {
	return uint32(len(points))
}
