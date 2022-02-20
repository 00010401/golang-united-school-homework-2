package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)

// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)
type sidesNumb int

const (
	SidesTriangle sidesNumb = 3
	SidesSquare   sidesNumb = 4
	SidesCircle   sidesNumb = 0
)

func CalcSquare(sideLen float64, sidesNum sidesNumb) (res float64) {
	if sidesNum == SidesTriangle {
		res = (math.Sqrt(3)) / 4 * sideLen * sideLen
	} else if sidesNum == SidesSquare {
		res = sideLen * sideLen
	} else if sidesNum == SidesCircle {
		res = (math.Pi) * sideLen * sideLen
	} else {
		res = 0
	}
	return
}
