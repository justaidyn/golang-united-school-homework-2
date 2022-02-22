package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
type caclType int

const SidesTriangle caclType = 3
const SidesSquare caclType = 4
const SidesCircle caclType = 0

func CalcSquare(sideLen float64, sidesNum caclType) float64 {
	switch sidesNum {
	case SidesTriangle:
		return ((math.Pow(sideLen, 2) * math.Pow(3, 0.5)) / 4)
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return (math.Pi * (math.Pow(sideLen, 2)))
	default:
		return 0
	}
}
