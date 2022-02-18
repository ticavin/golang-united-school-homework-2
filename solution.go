package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type SidesNumType int

var SidesTriangle SidesNumType = 3
var SidesSquare SidesNumType = 4
var SidesCircle SidesNumType = 0

func CalcSquare(sideLen float64, sidesNum SidesNumType) float64 {
	if sidesNum == 1 {
		return sideLen
	} else if sidesNum == 3 {
		return math.Sqrt(3) / 4 * sideLen * sideLen
	} else if sidesNum == 0 {
		return math.Pi * sideLen
	} else if sidesNum == 4 {
		return sideLen * sideLen
	} else {
		return 0
	}

}
