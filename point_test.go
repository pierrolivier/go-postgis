package postgis

import (
	"testing"
)

func TestPointBasic(t *testing.T) {
	// Test basic Point
	p := Point{X: 1.0, Y: 2.0}

	// Test Value() method
	value, err := p.Value()
	if err != nil {
		t.Errorf("Point.Value() failed: %v", err)
	}

	// Test Scan() method
	var p2 Point
	err = p2.Scan(value)
	if err != nil {
		t.Errorf("Point.Scan() failed: %v", err)
	}

	// Verify the data
	if p2.X != p.X || p2.Y != p.Y {
		t.Errorf("Point mismatch: expected (%f, %f), got (%f, %f)",
			p.X, p.Y, p2.X, p2.Y)
	}
}

func TestPointSBasic(t *testing.T) {
	// Test Point with SRID 4326
	p := PointS{SRID: 4326, X: -122.4194, Y: 37.7749}

	// Test Value() method
	value, err := p.Value()
	if err != nil {
		t.Errorf("PointS.Value() failed: %v", err)
	}

	// Test Scan() method
	var p2 PointS
	err = p2.Scan(value)
	if err != nil {
		t.Errorf("PointS.Scan() failed: %v", err)
	}

	// Verify SRID
	if p2.SRID != 4326 {
		t.Errorf("Expected SRID 4326, got %d", p2.SRID)
	}

	// Verify the data
	if p2.X != p.X || p2.Y != p.Y {
		t.Errorf("Point mismatch: expected (%f, %f), got (%f, %f)",
			p.X, p.Y, p2.X, p2.Y)
	}
}

func TestPointZSBasic(t *testing.T) {
	// Test Point with Z coordinate and SRID 4326
	p := PointZS{SRID: 4326, X: -122.4194, Y: 37.7749, Z: 100.0}

	// Test Value() method
	value, err := p.Value()
	if err != nil {
		t.Errorf("PointZS.Value() failed: %v", err)
	}

	// Test Scan() method
	var p2 PointZS
	err = p2.Scan(value)
	if err != nil {
		t.Errorf("PointZS.Scan() failed: %v", err)
	}

	// Verify SRID
	if p2.SRID != 4326 {
		t.Errorf("Expected SRID 4326, got %d", p2.SRID)
	}

	// Verify the data
	if p2.X != p.X || p2.Y != p.Y || p2.Z != p.Z {
		t.Errorf("Point mismatch: expected (%f, %f, %f), got (%f, %f, %f)",
			p.X, p.Y, p.Z, p2.X, p2.Y, p2.Z)
	}
}

func TestPointGetType(t *testing.T) {
	// Test GetType() methods for different Point variants
	tests := []struct {
		name     string
		geometry Geometry
		expected uint32
	}{
		{"Point", &Point{}, BuildWKBType(WKBPoint, CoordXY, false)},
		{"PointZ", &PointZ{}, BuildWKBType(WKBPoint, CoordXYZ, false)},
		{"PointM", &PointM{}, BuildWKBType(WKBPoint, CoordXYM, false)},
		{"PointZM", &PointZM{}, BuildWKBType(WKBPoint, CoordXYZM, false)},
		{"PointS", &PointS{}, BuildWKBType(WKBPoint, CoordXY, true)},
		{"PointZS", &PointZS{}, BuildWKBType(WKBPoint, CoordXYZ, true)},
		{"PointMS", &PointMS{}, BuildWKBType(WKBPoint, CoordXYM, true)},
		{"PointZMS", &PointZMS{}, BuildWKBType(WKBPoint, CoordXYZM, true)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.geometry.GetType(); got != test.expected {
				t.Errorf("%s.GetType() = 0x%X, expected 0x%X", test.name, got, test.expected)
			}
		})
	}
}
