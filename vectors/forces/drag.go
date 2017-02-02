package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

const defaultDragCoef = 0.02

// Drag force
type Drag struct {
	mover *vectors.Mover
	coef  float64
}

// CreateDrag creates a drag force
func CreateDrag(mover *vectors.Mover, coef float64) *Drag {
	// Fd = -1/2 * ro * v^2 * A * Cd * v(^)
	// Fd - Drag force vector
	// ro - density of the environment (air/liquid)
	// v - velocity of the object
	// A - frontal area
	// Cd - drag coefficient
	// v(^) - velocity unit vector
	//
	// Simplified: Fd = -1 * m.velocity.mag() ^ 2 * Cd * m.velocity.norm()

	var dragC float64

	if coef > 0 {
		dragC = coef
	} else {
		dragC = defaultDragCoef
	}

	return &Drag{mover, dragC}
}

// GetForce returns the force vector of gravity
func (d *Drag) GetForce() *vectors.Vector {
	force := d.mover.GetVelocity().Copy()
	speed := force.Magnitude()
	mag := -1 * speed * speed * d.coef

	force.Normalize()
	force.Multiply(mag)

	return force
}
