package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// Gravity force
type Gravity struct {
	Force
}

// CreateGravity creates a gravity force
func CreateGravity() *Gravity {
	// todo values
	return &Gravity{Force{vector: vectors.NewVector(0, 0.08)}}
}

// GetForce returns the force vector of gravity
func (g *Gravity) GetForce() *vectors.Vector {
	return g.vector
}
