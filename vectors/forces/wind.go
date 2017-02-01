package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// CreateWind creates a wind with
func CreateWind(force float64) *vectors.Vector {
	// todo
	return vectors.NewVector(0.5*force, 0)
}
