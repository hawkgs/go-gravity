package rendering

import (
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

	// Create group of rects
	rectGroup := []*sdl.Rect{
		&sdl.Rect{X: 100, Y: 100, W: 10, H: 10},
		&sdl.Rect{X: 100, Y: 100, W: 10, H: 10},
		&sdl.Rect{X: 100, Y: 100, W: 10, H: 10}}

	psConf := particles.NewConf(true, vectors.NewVector(100, 100), vectors.NewVector(WindowWidth, WindowHeight))
	ps := particles.NewParticleSystem(rectGroup, psConf)

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
