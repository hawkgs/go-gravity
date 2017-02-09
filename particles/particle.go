package particles

import (
	"math/rand"

	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

const particleLifespan = 100
const defaultPushDuration = 60 // Frames

var particleCt int

// Particle keeps vectors.Mover and adds a lifespan property
type Particle struct {
	mover    *vectors.Mover
	push     *forces.Push
	lifespan int
	id       int
}

// NewParticle creates an object of type Particle (constructor)
func NewParticle(mover *vectors.Mover) *Particle {
	particleCt++
	pushForce := getRandomPushDirection()
	return &Particle{mover, pushForce, particleLifespan, particleCt}
}

func getRandomPushDirection() *forces.Push {
	dir := float64(rand.Intn(359))
	return forces.CreatePush(dir, 0.5, defaultPushDuration)
}
