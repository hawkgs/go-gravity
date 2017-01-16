package vector

import vector "github.com/hAWKdv/go-gravity/vectors"

// CreateWind creates a wind with
func CreateWind(force float64) *vector.Vector {
	// todo
	return vector.NewVector(0.5*force, 0)
}
