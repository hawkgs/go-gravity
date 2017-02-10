package particles

// UpdateSystemSdl2 execures UpdateSdl2 on every mover/particle
func (ps *ParticleSystem) UpdateSystemSdl2(updateSdlObj func(obj interface{})) {
	ps.UpdateSystem(func(p *Particle) {
		p.mover.UpdateSdl2()
		updateSdlObj(p.mover.Obj)
	})
}
