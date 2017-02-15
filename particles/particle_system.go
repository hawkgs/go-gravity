package particles

import (
	"reflect"
	"sync"

	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

// Conf represents the configuration structure of a particle system
type Conf struct {
	continious       bool
	particleLifespan int
	location         *vectors.Vector
	container        *vectors.Vector
}

// ParticleSystem describes a set a of particles
type ParticleSystem struct {
	particles []*Particle
	conf      *Conf
	gravity   *forces.Gravity
	frame     int
}

// NewConf creates a configuration object for a particle system
func NewConf(continious bool, pLifespan int, location, container *vectors.Vector) *Conf {
	return &Conf{continious, pLifespan, location, container}
}

// NewParticleSystem creates an object of type ParticleSystem (constructor)
func NewParticleSystem(objs interface{}, conf *Conf) *ParticleSystem {
	var particles []*Particle

	if reflect.TypeOf(objs).Kind() == reflect.Slice {
		s := reflect.ValueOf(objs)

		for i := 0; i < s.Len(); i++ {
			mover := vectors.NewMover(s.Index(i).Interface(), conf.location.Copy(), conf.container)
			particles = append(particles, NewParticle(mover, conf.particleLifespan))
		}
	}

	return &ParticleSystem{particles, conf, forces.CreateGravity(), 0}
}

// UpdateSystem is a genetic method for updating particles in the system
func (ps *ParticleSystem) UpdateSystem(update func(p *Particle)) {
	wg := sync.WaitGroup{}

	for _, particle := range ps.particles {
		wg.Add(1)

		go func(particle *Particle) {
			if ps.conf.continious || particle.lifespan > 0 {
				ps.applyForces(particle)
				update(particle)
				particle.mover.BounceOff()

				particle.lifespan--
			} else if !ps.conf.continious && particle.lifespan == 0 {
				particle.Reset(ps.conf.particleLifespan)
			}
			wg.Done()
		}(particle)
	}

	ps.frame++
	wg.Wait()
}

func (ps *ParticleSystem) applyForces(p *Particle) {
	if ps.frame <= p.retardation {
		return
	}

	p.mover.ApplyForce(ps.gravity.GetForce())
	p.mover.ApplyForce(p.friction.GetForce())
	p.mover.ApplyForce(p.push.GetForce())
}
