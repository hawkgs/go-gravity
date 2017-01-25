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
func NewVector(x, y float64) *Vector {
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

// Limit vector's magnitude by a provided value
func (v *Vector) Limit(mag float64) {
	curr := v.Magnitude()

	if curr > mag && mag > 0 {
		ratio := curr / mag
		v.Divide(ratio)
	}
}

// Heading returns the heading of the vector in degrees
func (v *Vector) Heading() int {
	angle := math.Atan2(v.y, v.x)

	return RadToDeg(angle)
}

// Rotate vector by a degree angle
func (v *Vector) Rotate(deg int) {
	rads := DegToRad(deg)
	cos := math.Cos(rads)
	sin := math.Sin(rads)

	v.x = v.x*cos - v.y*sin
	v.y = v.x*sin + v.y*cos
}

// Normalize sets the magnitude to 1
func (v *Vector) Normalize() {
	mag := v.Magnitude()

	if mag != 0 {
		v.Divide(mag)
	}
}

// PixelLoc returns the rounded values of vector's X and Y
func (v *Vector) PixelLoc() (int, int) {
	return round(v.x), round(v.y)
}

// Copy creates a copy of the vector
func (v *Vector) Copy() *Vector {
	return NewVector(v.x, v.y)
}

func round(n float64) int {
	if n < 0 {
		return math.Ceil(n - 0.5)
	}
	return math.Floor(n + 0.5)
}
