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
func (v *Vector) Add(u *Vector) {
	v.x += u.x
	v.y += u.y
}

// Subtract performs subtraction of the current and provided as an argument vectors
func (v *Vector) Subtract(u *Vector) {
	v.x -= u.x
	v.y -= u.y
}

// Multiply performs multiplication of the current vector by N
func (v *Vector) Multiply(n float64) {
	v.x *= n
	v.y *= n
}

// Divide performs division of the current vector by N
func (v *Vector) Divide(n float64) {
	v.x /= n
	v.y /= n
}

// Magnitude returns the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

// Normalize sets the magnitude to 1
func (v *Vector) Normalize() {
	mag := v.Magnitude()

	if mag != 0 {
		v.Divide(mag)
	}
}

// Copy creates a copy of the vector
func (v *Vector) Copy() *Vector {
	return NewVector(v.x, v.y)
}
