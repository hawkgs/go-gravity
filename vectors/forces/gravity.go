package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// CreateGravity creates a gravity force
func CreateGravity() *vectors.Vector {
	// todo
	return vectors.NewVector(0, 0.3)
}
