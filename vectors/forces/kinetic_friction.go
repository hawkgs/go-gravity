package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// KineticFriction force
type KineticFriction struct {
	mover     *vectors.Mover
	magnitude float64
}

// CreateKineticFriction creates a kinetic friction force
func CreateKineticFriction(mover *vectors.Mover) *KineticFriction {
	// NOTE(Georgi): Currently hardcoded
	coef := 0.02
	normalForce := 1.0

	return &KineticFriction{mover, coef * normalForce}
}

// GetForce returns the force vector of kinetic friction
func (kf *KineticFriction) GetForce() *vectors.Vector {
	friction := kf.mover.GetVelocity().Copy()
	friction.Multiply(-1)
	friction.Normalize()
	friction.Multiply(kf.magnitude)

	return friction
}
