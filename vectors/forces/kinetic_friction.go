package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// KineticFriction force
type KineticFriction struct {
	Force
}

// CreateKineticFriction creates a kinetic friction force
func CreateKineticFriction() *KineticFriction {
	// todo
	return &KineticFriction{}
}

// GetForce returns the force vector of kinetic friction
func (kf *KineticFriction) GetForce() *vectors.Vector {
	return kf.vector
}
