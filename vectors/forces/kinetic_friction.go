package forces

import "github.com/hawkgs/go-gravity/vectors/vectors"

const (
	// NormalForce is the default normal force for the environment
	NormalForce = 1.0
	// FCoef is the friction coefficient
	FCoef = 0.02
)

// KineticFriction force
type KineticFriction struct {
	mover     *vectors.Mover
	magnitude float64
}

// CreateKineticFriction creates a kinetic friction force
func CreateKineticFriction(mover *vectors.Mover) *KineticFriction {
	return &KineticFriction{mover, FCoef * NormalForce}
}

// GetForce returns the force vector of kinetic friction
func (kf *KineticFriction) GetForce() *vectors.Vector {
	friction := kf.mover.GetVelocity().Copy()
	friction.Multiply(-1)
	friction.Normalize()
	friction.Multiply(kf.magnitude)

	return friction
}
