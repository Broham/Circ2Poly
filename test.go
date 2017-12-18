package main

import (
	"fmt"
	"github.com/broham/Circ2Poly/polygon"
	"strings"
)

var (
	geo      = ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingNotSymmetric)
	srcLat   = 29.8477935791016
	srcLong  = -97.9479217529297
	radius   = 2000.0
	sections = 10.0
)

func main() {
	points := polygon.ApproximateWKTCircle(srcLat, srcLong, radius, sections)
	fmt.Println("Points:", points)
}
