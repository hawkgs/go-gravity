package particles

// UpdateSystemSdl2 execures UpdateSdl2 on every mover/particle
func (ps *ParticleSystem) UpdateSystemSdl2() {
	ps.UpdateSystem(func(p *Particle) {
		p.mover.UpdateSdl2()
	})
}
