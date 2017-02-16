package particles

import (
	"math/rand"

	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

const defaultPushDuration = 15 // Frames

var particleCt int

// Particle keeps vectors.Mover and adds a lifespan property
type Particle struct {
	mover       *vectors.Mover
	push        *forces.Push
	friction    *forces.KineticFriction
	retardation int
	lifespan    int
	id          int
}

// NewParticle creates an object of type Particle (constructor)
func NewParticle(mover *vectors.Mover, lifespan int, maxRetardation int) *Particle {
	particleCt++
	pushForce := getRandomPushDirection()
	kinFriction := forces.CreateKineticFriction(mover)
	retardation := rand.Intn(maxRetardation)
	lifespan += retardation

	return &Particle{mover, pushForce, kinFriction, retardation, lifespan, particleCt}
}

// Reset a particle to its initial state
func (p *Particle) Reset(lifespan int) {
	p.lifespan = lifespan
	p.mover.ResetLocation()
	p.push = getRandomPushDirection()
}

func getRandomPushDirection() *forces.Push {
	dir := float64(rand.Intn(359))
	// todo: change magnitude
	return forces.CreatePush(dir, 0.2, defaultPushDuration)
}
