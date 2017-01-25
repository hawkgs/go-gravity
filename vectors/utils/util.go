package vector

import (
	"math"

	vector "github.com/hAWKdv/go-gravity/vectors"
)

// Distance calculates the Eucleadean distance between two vectors
func Distance(u *vector.GVector, v *vector.GVector) float64 {
	return math.Sqrt(math.Pow(u.x-v.x, 2) + math.Pow(u.y-v.y, 2))
}
