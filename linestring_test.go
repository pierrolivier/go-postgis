package postgis

import (
	"testing"
)

func TestLineString(t *testing.T) {
	// Test basic LineString
	ls := LineString{
		Points: []Point{
			{X: 1.0, Y: 2.0},
			{X: 3.0, Y: 4.0},
			{X: 5.0, Y: 6.0},
		},
	}

	// Test Value() method
	value, err := ls.Value()
	if err != nil {
		t.Errorf("LineString.Value() failed: %v", err)
	}

	// Test Scan() method
	var ls2 LineString
	err = ls2.Scan(value)
	if err != nil {
		t.Errorf("LineString.Scan() failed: %v", err)
	}

	// Verify the data
	if len(ls2.Points) != 3 {
		t.Errorf("Expected 3 points, got %d", len(ls2.Points))
	}

	for i, point := range ls2.Points {
		if point.X != ls.Points[i].X || point.Y != ls.Points[i].Y {
			t.Errorf("Point %d mismatch: expected (%f, %f), got (%f, %f)",
				i, ls.Points[i].X, ls.Points[i].Y, point.X, point.Y)
		}
	}
}

func TestLineStringS_4326(t *testing.T) {
	// Test LineString with SRID 4326 (WGS84)
	ls := LineStringS{
		SRID: 4326,
		Points: []Point{
			{X: -122.4194, Y: 37.7749}, // San Francisco
			{X: -74.0060, Y: 40.7128},  // New York
			{X: 2.3522, Y: 48.8566},    // Paris
		},
	}

	// Test Value() method
	value, err := ls.Value()
	if err != nil {
		t.Errorf("LineStringS.Value() failed: %v", err)
	}

	// Test Scan() method
	var ls2 LineStringS
	err = ls2.Scan(value)
	if err != nil {
		t.Errorf("LineStringS.Scan() failed: %v", err)
	}

	// Verify SRID
	if ls2.SRID != 4326 {
		t.Errorf("Expected SRID 4326, got %d", ls2.SRID)
	}

	// Verify the data
	if len(ls2.Points) != 3 {
		t.Errorf("Expected 3 points, got %d", len(ls2.Points))
	}

	for i, point := range ls2.Points {
		if point.X != ls.Points[i].X || point.Y != ls.Points[i].Y {
			t.Errorf("Point %d mismatch: expected (%f, %f), got (%f, %f)",
				i, ls.Points[i].X, ls.Points[i].Y, point.X, point.Y)
		}
	}
}

func TestLineStringZ(t *testing.T) {
	// Test LineString with Z coordinates
	ls := LineStringZ{
		Points: []PointZ{
			{X: 1.0, Y: 2.0, Z: 10.0},
			{X: 3.0, Y: 4.0, Z: 20.0},
			{X: 5.0, Y: 6.0, Z: 30.0},
		},
	}

	// Test Value() method
	value, err := ls.Value()
	if err != nil {
		t.Errorf("LineStringZ.Value() failed: %v", err)
	}

	// Test Scan() method
	var ls2 LineStringZ
	err = ls2.Scan(value)
	if err != nil {
		t.Errorf("LineStringZ.Scan() failed: %v", err)
	}

	// Verify the data
	if len(ls2.Points) != 3 {
		t.Errorf("Expected 3 points, got %d", len(ls2.Points))
	}

	for i, point := range ls2.Points {
		if point.X != ls.Points[i].X || point.Y != ls.Points[i].Y || point.Z != ls.Points[i].Z {
			t.Errorf("Point %d mismatch: expected (%f, %f, %f), got (%f, %f, %f)",
				i, ls.Points[i].X, ls.Points[i].Y, ls.Points[i].Z, point.X, point.Y, point.Z)
		}
	}
}

func TestLineStringZS_4326(t *testing.T) {
	// Test LineString with Z coordinates and SRID 4326
	ls := LineStringZS{
		SRID: 4326,
		Points: []PointZ{
			{X: -122.4194, Y: 37.7749, Z: 100.0}, // San Francisco with elevation
			{X: -74.0060, Y: 40.7128, Z: 50.0},   // New York with elevation
			{X: 2.3522, Y: 48.8566, Z: 75.0},     // Paris with elevation
		},
	}

	// Test Value() method
	value, err := ls.Value()
	if err != nil {
		t.Errorf("LineStringZS.Value() failed: %v", err)
	}

	// Test Scan() method
	var ls2 LineStringZS
	err = ls2.Scan(value)
	if err != nil {
		t.Errorf("LineStringZS.Scan() failed: %v", err)
	}

	// Verify SRID
	if ls2.SRID != 4326 {
		t.Errorf("Expected SRID 4326, got %d", ls2.SRID)
	}

	// Verify the data
	if len(ls2.Points) != 3 {
		t.Errorf("Expected 3 points, got %d", len(ls2.Points))
	}

	for i, point := range ls2.Points {
		if point.X != ls.Points[i].X || point.Y != ls.Points[i].Y || point.Z != ls.Points[i].Z {
			t.Errorf("Point %d mismatch: expected (%f, %f, %f), got (%f, %f, %f)",
				i, ls.Points[i].X, ls.Points[i].Y, ls.Points[i].Z, point.X, point.Y, point.Z)
		}
	}
}

func TestLineStringGetType(t *testing.T) {
	// Test GetType() methods for different LineString variants
	tests := []struct {
		name     string
		geometry Geometry
		expected uint32
	}{
		{"LineString", &LineString{}, 2},
		{"LineStringZ", &LineStringZ{}, 0x80000002},
		{"LineStringM", &LineStringM{}, 0x40000002},
		{"LineStringZM", &LineStringZM{}, 0xC0000002},
		{"LineStringS", &LineStringS{}, 0x20000002},
		{"LineStringZS", &LineStringZS{}, 0xA0000002},
		{"LineStringMS", &LineStringMS{}, 0x60000002},
		{"LineStringZMS", &LineStringZMS{}, 0xE0000002},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.geometry.GetType(); got != test.expected {
				t.Errorf("%s.GetType() = 0x%X, expected 0x%X", test.name, got, test.expected)
			}
		})
	}
}

func TestEmptyLineString(t *testing.T) {
	// Test empty LineString
	ls := LineString{Points: []Point{}}

	// Test Value() method
	value, err := ls.Value()
	if err != nil {
		t.Errorf("Empty LineString.Value() failed: %v", err)
	}

	// Test Scan() method
	var ls2 LineString
	err = ls2.Scan(value)
	if err != nil {
		t.Errorf("Empty LineString.Scan() failed: %v", err)
	}

	// Verify the data
	if len(ls2.Points) != 0 {
		t.Errorf("Expected 0 points, got %d", len(ls2.Points))
	}
}
