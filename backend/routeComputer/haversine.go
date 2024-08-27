package routecomputer

import "math"

const earthRadius float64 = 6371

func haversine(lat1, long1, lat2, long2 float64) float64 {
	// distances
	dLat := (lat2 - lat1) * math.Pi / 180.0
	dLong := (long2 - long1) * math.Pi / 180.0
	// convert to radians
	lat1Rad := lat1 * math.Pi / 180.0
	lat2Rad := lat2 * math.Pi / 180.0

	a := math.Pow(math.Sin(dLat / 2), 2) + math.Pow(math.Sin(dLong / 2), 2) * math.Cos(lat1Rad) * math.Cos(lat2Rad)
	c := 2 * math.Asin(math.Sqrt(a))
	return earthRadius * c
}
