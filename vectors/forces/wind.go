package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// Wind force
type Wind struct {
	Force
}

// CreateWind creates a wind force
func CreateWind() *Wind {
	// todo values
	return &Wind{Force{vectors.NewVector(0.01, 0)}}
}

// GetForce returns the force vector of gravity
func (w *Wind) GetForce() *vectors.Vector {
	return w.vector
}
