package routecomputer

import (
	"math"
	"testing"
)

// TestHaversine tests the Haversine function for calculating distances between two points on the Earth's surface.
//
// ** ChatGPT Generated test **
func TestHaversine(t *testing.T) {
	tests := []struct {
		lat1, long1, lat2, long2 float64
		expectedDistance         float64
	}{
		{lat1: 0, long1: 0, lat2: 0, long2: 0, expectedDistance: 0},                                  // Same point
		{lat1: 36.12, long1: -86.67, lat2: 33.94, long2: -118.40, expectedDistance: 2887.26},         // Example: Nashville, TN to Los Angeles, CA
		{lat1: 51.5, long1: 0, lat2: 48.8566, long2: 2.3522, expectedDistance: 338.77},               // London to Paris
		{lat1: 40.7128, long1: -74.0060, lat2: 34.0522, long2: -118.2437, expectedDistance: 3935.95}, // New York to Los Angeles
	}

	for _, tt := range tests {
		distance := haversine(tt.lat1, tt.long1, tt.lat2, tt.long2)
		if math.Abs(distance-tt.expectedDistance) > 1 {
			t.Errorf("haversine(%f, %f, %f, %f) = %f; want %f", tt.lat1, tt.long1, tt.lat2, tt.long2, distance, tt.expectedDistance)
		}
	}
}
