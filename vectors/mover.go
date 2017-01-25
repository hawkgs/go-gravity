package vector

import (
	"math"

	vector "github.com/hAWKdv/go-gravity/vectors"
)

// Mover describes a basic moveable object/particle
type Mover struct {
	obj          interface{}
	acceleration *vector.GVector
	velocity     *vector.GVector
	location     *vector.GVector
	mass         float64
}

// NewMover creates an object of type Mover (constructor)
func NewMover(obj interface{}, location *vector.GVector) *Mover {
	var mover *Mover

	if obj != nil {
		mover = &Mover{obj: obj}
	} else {
		mover = &Mover{}
	}

	mover.acceleration = NewVector(0, 0)
	mover.velocity = NewVector(0, 0)

	if location != nil {
		mover.location = location
	} else {
		mover.location = NewVector(0, 0)
	}

	return mover
}

// SetMass assigns the mass argument to the object's mass
func (m *Mover) SetMass(mass float64) {
	m.mass = mass
}

// ApplyForce adds the force vector the object's acceleration vector
func (m *Mover) ApplyForce(force *vector.Vector) {
	// Newton's 2nd law: Acceleration = Sum of all forces / Mass
	fCopy := force.Copy()
	// Apply the full law only if the mass has been set
	if mass > 0 {
		fCopy.Divide(mass)
	}
	m.acceleration.Add(fCopy)
}

// Update modifies the object's location depending on the applied forces;
// Should be called on every rendering iteration
func (m *Mover) Update() {
	// We keep the velocity only for correctness based on physics laws
	m.velocity.Add(m.acceleration)
	m.location.Add(m.velocity)
	// Clear the acceleration
	m.acceleration.Multiply(0)
}

// PixelLoc returns the rounded values of location's X and Y which are ready for rendering
func (m *Mover) PixelLoc() (int, int) {
	return round(m.location.x), round(m.location.y)
}

func round(n float64) int {
	if n < 0 {
		return math.Ceil(n - 0.5)
	}
	return math.Floor(n + 0.5)
}
