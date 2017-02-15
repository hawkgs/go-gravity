package vectors

import "github.com/hAWKdv/go-gravity/vectors/utils"

// MoverDefMass keeps the default mass of the Mover
const MoverDefMass = 1

// Mover describes a basic moveable object/particle
type Mover struct {
	Obj           interface{}
	acceleration  *Vector
	velocity      *Vector
	location      *Vector
	container     *Vector
	initLoc       *Vector
	mass          float64
	limit         float64
	aAcceleration float64
	aVelocity     float64
	angle         float64
}

// NewMover creates an object of type Mover (constructor)
func NewMover(obj interface{}, location *Vector, container *Vector) *Mover {
	var mover *Mover

	if obj != nil {
		mover = &Mover{Obj: obj}
	} else {
		mover = &Mover{}
	}

	mover.acceleration = NewVector(0, 0)
	mover.velocity = NewVector(0, 0)
	mover.mass = MoverDefMass
	mover.limit = 0
	mover.container = container
	mover.location = location
	mover.initLoc = location.Copy()

	return mover
}

// SetMass assigns the mass argument to the object's mass
func (m *Mover) SetMass(mass float64) {
	if mass < 1 {
		return
	}
	m.mass = mass
}

// SetLimit puts a velocity limit when accelerating
func (m *Mover) SetLimit(limit float64) {
	m.limit = limit
}

// GetVelocity return the current velocity vector of the Mover
func (m *Mover) GetVelocity() *Vector {
	return m.velocity
}

// ApplyForce adds the force vector the object's acceleration vector
func (m *Mover) ApplyForce(force *Vector) {
	if force == nil {
		return
	}

	// Newton's 2nd law: Acceleration = Sum of all forces / Mass
	fCopy := force.Copy()
	fCopy.Divide(m.mass)
	m.acceleration.Add(fCopy)
}

// Update modifies the object's location depending on the applied forces;
// Should be called on every rendering iteration
func (m *Mover) Update() {
	// We keep the velocity only for correctness based on physics laws
	m.velocity.Add(m.acceleration)

	// Apply velocity limit, if there is any.
	if m.limit > 0 {
		m.velocity.Limit(m.limit)
	}

	m.location.Add(m.velocity)
	// Clear the acceleration
	m.acceleration.Multiply(0)
}

// PixelLoc returns the rounded values of location's X and Y which are ready for rendering
func (m *Mover) PixelLoc() (int, int) {
	return utils.Round(m.location.X), utils.Round(m.location.Y)
}

// BounceOff keeps the mover within its container (bounces off) when it reaches an edge
func (m *Mover) BounceOff() {
	if m.container == nil {
		return
	}

	// Right/left
	if m.location.X > m.container.X {
		m.velocity.X *= -1
		m.location.X = m.container.X
	} else if m.location.X < 0 {
		m.velocity.X *= -1
		m.location.X = 0
	}

	// Bottom/top
	if m.location.Y > m.container.Y {
		m.velocity.Y *= -1
		m.location.Y = m.container.Y
	} else if m.location.Y < 0 {
		m.velocity.Y *= -1
		m.location.Y = 0
	}
}

// ResetLocation sets location vector to its initial value
func (m *Mover) ResetLocation() {
	m.location = m.initLoc.Copy()
}
