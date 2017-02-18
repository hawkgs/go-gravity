package utils

import "math"

// Distance calculates the Eucleadean distance between two vectors
// func Distance(u *vector.Vector, v *vector.Vector) float64 {
// 	return math.Sqrt(math.Pow(u.x-v.x, 2) + math.Pow(u.y-v.y, 2))
// }

// RadToDeg converts radians to degrees
func RadToDeg(rad float64) float64 {
	return rad * 180 / math.Pi
}

// DegToRad converts degrees to radians
func DegToRad(deg float64) float64 {
	return 2 * math.Pi * (deg / 360)
}

// Round performs rounding of a float type to the nearest int
func Round(n float64) int {
	if n > 0 {
		return int(math.Floor(n + 0.5))
	}
	return int(math.Ceil(n - 0.5))
}
