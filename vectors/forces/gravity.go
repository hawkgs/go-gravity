package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

const gravitationalPull = 0.08

// Gravity force
type Gravity struct {
	Force
}

// CreateGravity creates a gravity force
func CreateGravity() *Gravity {
	return &Gravity{Force{vector: vectors.NewVector(0, gravitationalPull)}}
}

// GetForce returns the force vector of gravity
func (g *Gravity) GetForce() *vectors.Vector {
	return g.vector
}
