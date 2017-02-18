package rendering

import (
	"math/rand"

	"github.com/hAWKdv/go-gravity/particles"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
	"github.com/veandco/go-sdl2/sdl"
)

//RenderParticleSystem used for rendering a particle system in SDL2
func RenderParticleSystem() int {
	window, exit := CreateSdlWindow()
	if exit != 0 {
		return exit
	}
	defer window.Destroy()

	// Create renderer
	renderer, exit := CreateSdlRenderer(window)
	if exit != 0 {
		return exit
	}
	defer renderer.Destroy()
	renderer.Clear()

	// Create masses
	var masses []float64
	for i := 0; i < 100; i++ {
		m := rand.Float64() + 1.0
		masses = append(masses, m)
	}

	// Create group of rects
	var rectGroup []*sdl.Rect
	for i := 0; i < 100; i++ {
		m := int32(masses[i]*10) - 5
		rectGroup = append(rectGroup, &sdl.Rect{X: 100, Y: 100, W: m, H: m})
		// rectGroup = append(rectGroup, &sdl.Rect{X: 100, Y: 100, W: 10, H: 10})
	}

	psConf := particles.NewConf(false, 100, 15, vectors.NewVector(100, 100), vectors.NewVector(WindowWidth, WindowHeight))

	ps := particles.NewParticleSystem(rectGroup, psConf)
	ps.SetMasses(masses)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
			}
		}

		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.FillRect(&sdl.Rect{X: 0, Y: 0, W: WindowWidth, H: WindowHeight})

		ps.UpdateSystemSdl2(func(obj interface{}) {
			renderer.SetDrawColor(0, 0, 0, 255)
			switch obj.(type) {
			case *sdl.Rect:
				renderer.DrawRect(obj.(*sdl.Rect))
				break
			}
		})

		renderer.Present()
		sdl.Delay(1000 / FrameRate)
	}

	return 0
}
