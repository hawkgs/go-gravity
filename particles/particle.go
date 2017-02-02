package particles

import "github.com/hAWKdv/go-gravity/vectors/vectors"

const particleLifespan = 100

// Particle keeps vectors.Mover and adds a lifespan property
type Particle struct {
	mover    *vectors.Mover
	lifespan int
}

// NewParticle creates an object of type Particle (constructor)
func NewParticle(mover *vectors.Mover) *Particle {
	return &Particle{mover, particleLifespan}
}
