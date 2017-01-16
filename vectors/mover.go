package vector

import vector "github.com/hAWKdv/go-gravity/vectors"

// Mover describes a basic moveable object/particle
type Mover struct {
	vector.Vector
	acceleration vector.GVector
	velocity     vector.GVector
	location     vector.GVector
}

// ApplyForce ...
func (m *Mover) ApplyForce(force *vector.Vector) {
	// todo
}
