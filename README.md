# Circ2Poly
Accepts a lat/long, radius and number of sections and returns a GEOJSON polygon approximating that circle.  The higher the number of sections that you specify the closer the generated polygon will be to a circle.  Currently there is only one function - `ApproximateWKTCircle`

### Parameters
#### lat/long
The center of the circle that will be aproximated in decimal format

#### radius
The radius of the circle to approximate in meters

#### sections
The number of sides the shape generated will have.  The larger this number is the closer to an actual circle the polygon will approximate

#### Example
```
import (
	"fmt"
	"github.com/broham/Circ2Poly/polygon"
)

var (
	lat   = 29.8477935791016
	long  = -97.9479217529297
	radius   = 2000.0
	sections = 10.0
)

func main() {
	points := polygon.ApproximateWKTCircle(lat, long, radius, sections)
	fmt.Println("Points:", points)
}
```
generates the following polygon:

![Sample polygon](https://raw.githubusercontent.com/broham/Circ2Poly/master/sample.png)
