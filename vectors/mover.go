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

// ApplyForce ...
func (m *Mover) ApplyForce(force *vector.Vector) {
	// todo
}

// Update ...
func (m *Mover) Update() {
	// todo
}

// PixelLoc returns the rounded values of vector's X and Y
func (m *Mover) PixelLoc() (int, int) {
	return round(m.location.x), round(m.location.y)
}

func round(n float64) int {
	if n < 0 {
		return math.Ceil(n - 0.5)
	}
	return math.Floor(n + 0.5)
}
