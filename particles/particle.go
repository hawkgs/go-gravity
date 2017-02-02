package particles

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// Particle ...
type Particle struct {
	vectors.Mover
	lifespan int
}

// NewParticle ...
func NewParticle(obj interface{}, location *vectors.Vector, container *vectors.Vector) *Particle {
	return &Particle{vectors.Mover{Obj: obj}, 100}
}
