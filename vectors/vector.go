package vector

import "math"

// GVector represents a geometric/Eucleadean vector
type GVector interface {
	Add()
	Subtract()
	Multiply()
	Mangnitude()
	Normalize()
	Copy()
}

// Vector implements a GVector
type Vector struct {
	x float64
	y float64
}

// NewVector creates a new object of type Vector (basically constructor)
func NewVector(x float64, y float64) *Vector {
	return &Vector{x, y}
}

// Add performs addition of the current and provided as an argument vectors
func (v *Vector) Add(adder *Vector) {
	// todo
}

// Subtract performs subtraction of the current and provided as an argument vectors
func (v *Vector) Subtract(sub *Vector) {
	// todo
}

// Multiply performs multiplication of the current and provided as an argument vectors
func (v *Vector) Multiply(mul *Vector) {
	// todo
}

// Magnitude returns the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Normalize sets the magnitude to 1
func (v *Vector) Normalize() {
	// todo
}

// Copy creates a copy of the vector
func (v *Vector) Copy() *Vector {
	return NewVector(v.x, v.y)
}
