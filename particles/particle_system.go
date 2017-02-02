package particles

import "github.com/hAWKdv/go-gravity/vectors/vectors"

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
	// gravity   forces.Force
}

// NewParticleSystem creates an object of type ParticleSystem (constructor)
func NewParticleSystem(objs []interface{}, conf *Conf) *ParticleSystem {
	var particles []*Particle
	for _, obj := range objs {
		mover := vectors.NewMover(obj, conf.location, conf.container)
		particles = append(particles, NewParticle(mover))
	}

	return &ParticleSystem{particles, conf}
}

func (ps *ParticleSystem) UpdateSystem() {
	// todo
}
