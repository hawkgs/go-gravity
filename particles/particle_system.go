package particles

import (
	"reflect"

	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

// Conf represents the configuration structure of a particle system
type Conf struct {
	continious bool
	location   *vectors.Vector
	container  *vectors.Vector
}

// ParticleSystem describes a set a of particles
type ParticleSystem struct {
	particles []*Particle
	conf      *Conf
	gravity   *forces.Gravity
}

// NewConf creates a configuration object for a particle system
func NewConf(continious bool, location, container *vectors.Vector) *Conf {
	return &Conf{continious, location, container}
}

// NewParticleSystem creates an object of type ParticleSystem (constructor)
func NewParticleSystem(objs interface{}, conf *Conf) *ParticleSystem {
	var particles []*Particle

	if reflect.TypeOf(objs).Kind() == reflect.Slice {
		s := reflect.ValueOf(objs)

		for i := 0; i < s.Len(); i++ {
			mover := vectors.NewMover(s.Index(i).Interface(), conf.location.Copy(), conf.container)
			particles = append(particles, NewParticle(mover))
		}
	}

	return &ParticleSystem{particles, conf, forces.CreateGravity()}
}

// UpdateSystem is a genetic method for updating particles in the system
func (ps *ParticleSystem) UpdateSystem(update func(p *Particle)) {
	// wg := sync.WaitGroup{}

	for _, particle := range ps.particles {
		// wg.Add(1)

		// go func(particle *Particle) {
		ps.applyForces(particle)
		// todo run in go routine
		update(particle)
		particle.mover.BounceOff()

		if !ps.conf.continious {
			if particle.lifespan > 0 {
				particle.lifespan--
			} else {
				// to do remove particle
			}
		}
		// wg.Done()
		// }(particle)
	}

	// wg.Wait()
}

func (ps *ParticleSystem) applyForces(p *Particle) {
	p.mover.ApplyForce(ps.gravity.GetForce())
	p.mover.ApplyForce(p.friction.GetForce())
	p.mover.ApplyForce(p.push.GetForce())
}
