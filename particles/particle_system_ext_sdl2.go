package particles

import "github.com/veandco/go-sdl2/sdl"

// UpdateSystemSdl2 execures UpdateSdl2 on every mover/particle
func (ps *ParticleSystem) UpdateSystemSdl2(updateSdlObj func(obj interface{})) {
	ps.UpdateSystem(func(p *Particle) {
		p.mover.UpdateSdl2()

		sdl.Do(func() {
			updateSdlObj(p.mover.Obj)
		})
	})
}
