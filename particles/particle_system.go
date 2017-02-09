package particles

import (
	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

// Conf represents the configuration structure of a particle system
type Conf struct {
	continious  bool
	particleNum int
	location    *vectors.Vector
	container   *vectors.Vector
}

// ParticleSystem describes a set a of particles
type ParticleSystem struct {
	particles []*Particle
	conf      *Conf
	gravity   *forces.Gravity
}

// NewParticleSystem creates an object of type ParticleSystem (constructor)
func NewParticleSystem(objs []interface{}, conf *Conf) *ParticleSystem {
	var particles []*Particle
	for _, obj := range objs {
		mover := vectors.NewMover(obj, conf.location, conf.container)
		particles = append(particles, NewParticle(mover))
	}

	return &ParticleSystem{particles, conf, forces.CreateGravity()}
}

// UpdateSystem is a genetic method for updating particles in the system
func (ps *ParticleSystem) UpdateSystem(update func(p *Particle)) {
	for _, particle := range ps.particles {
		if ps.conf.continious {
			ps.applyForces(particle)
			update(particle)
		} else {
			if particle.lifespan > 0 {
				ps.applyForces(particle)
				update(particle)
				particle.lifespan--
			} else {
				// remove particle
			}
		}
	}
}

func (ps *ParticleSystem) applyForces(p *Particle) {
	p.mover.ApplyForce(ps.gravity.GetForce())
	p.mover.ApplyForce(p.friction.GetForce())
	p.mover.ApplyForce(p.push.GetForce())
}
