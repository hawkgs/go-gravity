package vectors

import (
	"math"

	"github.com/hAWKdv/go-gravity/vectors/utils"
)

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
	X float64
	Y float64
}

// NewVector creates a new object of type Vector (basically constructor)
func NewVector(x, y float64) *Vector {
	return &Vector{x, y}
}

// Add performs addition of the current and provided as an argument vectors
func (v *Vector) Add(u *Vector) {
	v.X += u.X
	v.Y += u.Y
}

// Subtract performs subtraction of the current and provided as an argument vectors
func (v *Vector) Subtract(u *Vector) {
	v.X -= u.X
	v.Y -= u.Y
}

// Multiply performs multiplication of the current vector by N
func (v *Vector) Multiply(n float64) {
	v.X *= n
	v.Y *= n
}

// Divide performs division of the current vector by N
func (v *Vector) Divide(n float64) {
	v.X /= n
	v.Y /= n
}

// Magnitude returns the magnitude of the vector
func (v *Vector) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Limit vector's magnitude by a provided value
func (v *Vector) Limit(mag float64) {
	curr := v.Magnitude()

	if curr > mag && mag > 0 {
		ratio := curr / mag
		v.Divide(ratio)
	}
}

// Heading returns the heading of the vector in degrees
func (v *Vector) Heading() (angle float64) {
	angle = math.Atan2(v.Y, v.X)
	angle = utils.RadToDeg(angle)

	if angle < 0 {
		angle += 360
	}

	return
}

// Rotate vector by a degree angle
func (v *Vector) Rotate(deg float64) {
	rads := utils.DegToRad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)

	v.X = v.X*cos - v.Y*sin
	v.Y = v.X*sin + v.Y*cos
}

// Normalize sets the magnitude to 1
func (v *Vector) Normalize() {
	mag := v.Magnitude()

	if mag != 0 {
		v.Divide(mag)
	}
}

// Distance calculates the Eucleadean distance between two vectors
func (v *Vector) Distance(u *Vector) float64 {
	return math.Sqrt(math.Pow(u.X-v.X, 2) + math.Pow(u.Y-v.Y, 2))
}

// Copy creates a copy of the vector
func (v *Vector) Copy() *Vector {
	return NewVector(v.X, v.Y)
}
