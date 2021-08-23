package utils

import (
	"math"

	"gotest.com/nodes-api/model"
)

type LocationUtils interface {
	Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64
}

/*
	This function calculate the distance between two points
*/
func Distance(location1 model.Location, location2 model.Location) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * location1.Lat / 180)
	radlat2 := float64(PI * location2.Lat / 180)

	theta := float64(location1.Lng - location2.Lng)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	//Distance in kilometers
	dist = dist * 1.609344

	return dist
}
