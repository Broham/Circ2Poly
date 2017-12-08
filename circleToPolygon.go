package circleToPolygon

import (
	"fmt"
	"math"
	"strings"
)

// function toRadians(angleInDegrees) {
//   return angleInDegrees * Math.PI / 180;
// }

// function toDegrees(angleInRadians) {
//   return angleInRadians * 180 / Math.PI;
// }

// function offset(c1, distance, bearing) {
//   var lat1 = toRadians(c1[1]);
//   var lon1 = toRadians(c1[0]);
//   var dByR = distance / 6378137; // distance divided by 6378137 (radius of the earth) wgs84
//   var lat = Math.asin(
//     Math.sin(lat1) * Math.cos(dByR) +
//     Math.cos(lat1) * Math.sin(dByR) * Math.cos(bearing));
//   var lon = lon1 + Math.atan2(
//       Math.sin(bearing) * Math.sin(dByR) * Math.cos(lat1),
//       Math.cos(dByR) - Math.sin(lat1) * Math.sin(lat));
//   return [toDegrees(lon), toDegrees(lat)];
// }

// module.exports = function circleToPolygon(center, radius, numberOfSegments) {
//   var n = numberOfSegments ? numberOfSegments : 32;
//   var flatCoordinates = [];
//   var coordinates = [];
//   for (let i = 0; i < n; ++i) {
//     flatCoordinates.push.apply(flatCoordinates, offset(center, radius, 2 * Math.PI * i / n));
//   }
//   flatCoordinates.push(flatCoordinates[0], flatCoordinates[1]);

//   for (let i = 0, j = 0; j < flatCoordinates.length; j += 2) {
//     coordinates[i++] = flatCoordinates.slice(j, j + 2);
//   }

//   return {
//     type: 'Polygon',
//     coordinates: [coordinates.reverse()]
//   };
// };

func GetPoint(lat float64, long float64, radius int, adjustment float64) string {
	var rad = float64(radius / 6378137.0)
	var retLat = math.Asin(math.Sin(lat)*math.Cos(rad) + math.Cos(lat)*math.Sin(rad)*math.Cos(adjustment))
	var retLong = math.Atan2(math.Sin(adjustment)*math.Sin(rad)*math.Cos(lat), math.Cos(rad)-math.Sin(lat)*math.Sin(lat))
	return fmt.Sprintf("[%v %v]", retLat, retLong)
}

func CircleToPolygon(lat float64, long float64, radius int, sections int) string {
	coordinates := make([]string, sections)
	for i := 1.0; i <= float64(sections); i++ {
		coordinates = append(coordinates, GetPoint(lat, long, radius, float64(2*math.Pi*i/float64(sections))))
	}
	fmt.Println("coordinates", coordinates)
	return strings.Join(coordinates, ",")
}
