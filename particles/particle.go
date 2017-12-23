package particles

import (
	"math/rand"

	"github.com/hawkgs/go-gravity/vectors/forces"
	"github.com/hawkgs/go-gravity/vectors/vectors"
)

const (
	defaultPushDuration  = 15 // Frames
	defaultPushMagnitude = 0.2
)

// Particle keeps vectors.Mover and adds a lifespan property
type Particle struct {
	mover         *vectors.Mover
	push          *forces.Push
	friction      *forces.KineticFriction
	retardation   int
	lifespan      int
	pushDuration  int
	pushMagnitude float64
}

// NewParticle creates an object of type Particle (constructor)
func NewParticle(mover *vectors.Mover, conf *Conf) *Particle {
	pushForce := getRandomPushDirection(conf.pushMagnitude, conf.pushDuration)
	kinFriction := forces.CreateKineticFriction(mover)
	retardation := rand.Intn(conf.maxRetardation)
	lifespan := conf.particleLifespan + retardation

	return &Particle{
		mover,
		pushForce,
		kinFriction,
		retardation,
		lifespan,
		conf.pushDuration,
		conf.pushMagnitude}
}

// Reset a particle to its initial state
func (p *Particle) Reset(lifespan int) {
	p.lifespan = lifespan
	p.mover.Reset()
	p.push = getRandomPushDirection(p.pushMagnitude, p.pushDuration)
}

func getRandomPushDirection(magnitude float64, duration int) *forces.Push {
	if magnitude <= 0 {
		magnitude = defaultPushMagnitude
	}
	if duration <= 0 {
		duration = defaultPushDuration
	}

	dir := float64(rand.Intn(359))
	// todo: change magnitude
	return forces.CreatePush(dir, magnitude, duration)
}
