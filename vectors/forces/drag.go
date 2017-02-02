package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// Drag force
type Drag struct {
	mover *vectors.Mover
	coef  float64
}

// CreateDrag creates a drag force
func CreateDrag(mover *vectors.Mover, coef float64) {
	// Fd = -1/2 * ro * v^2 * A * Cd * v(^)
	// Fd - Drag force vector
	// ro - density of the environment (air/liquid)
	// v - velocity of the object
	// A - frontal area
	// Cd - drag coefficient
	// v(^) - velocity unit vector
	//
	// Simplified: Fd = -1 * m.velocity.mag() ^ 2 * Cd * m.velocity.norm()
	return &Drag{mover}
}

func (d *Drag) GetForce() *vectors.Vector {
	force := mover.GetVelocity().Copy()
	speed := force.Magnitude()
	direction.Normalize()
	direction.Multiply(-1)

	vUnit = mover.GET
}
