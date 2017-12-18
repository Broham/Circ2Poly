package polygon

import (
	"fmt"
	"github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
	"strings"
)

func ApproximateWKTCircle(srcLat float64, srcLong float64, radius float64, sections float64) string {
	geo := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingNotSymmetric)
	points := make([]string, int(0))
	for i := 0.0; i <= sections; i++ {
		adjust := i / sections
		lat, long := geo.At(srcLat, srcLong, radius, adjust*360.0)
		points = append(points, fmt.Sprintf("%v %v", long, lat))
	}
	return fmt.Sprintf("POLYGON((%v))", strings.Join(points, ","))
}
