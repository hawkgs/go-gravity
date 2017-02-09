package particles

import "github.com/hAWKdv/go-gravity/vectors/vectors"

const particleLifespan = 100

var particleCt int

// Particle keeps vectors.Mover and adds a lifespan property
type Particle struct {
	mover    *vectors.Mover
	lifespan int
	id       int
}

// NewParticle creates an object of type Particle (constructor)
func NewParticle(mover *vectors.Mover) *Particle {
	particleCt++
	return &Particle{mover, particleLifespan, particleCt}
}
