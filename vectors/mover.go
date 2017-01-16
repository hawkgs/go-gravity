package vector

// Mover describes a basic moveable object/particle
type Mover struct {
	Vector
	acceleration Vector
	velocity     Vector
	location     Vector
}
