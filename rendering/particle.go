package rendering

import (
	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
	"github.com/veandco/go-sdl2/sdl"
)

// RenderParticle used for rendering a single particle in SDL2
func RenderParticle() int {
	// Create window
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

	// Create rect
	rect := &sdl.Rect{X: 100, Y: 100, W: 10, H: 10}
	mover := vectors.NewMover(
		rect,
		vectors.NewVector(float64(rect.X), float64(rect.Y)),
		vectors.NewVector(WindowWidth, WindowHeight))
	// mover.SetLimit(0.2)

	// Create some forces
	gravity := forces.CreateGravity()
	wind := forces.CreateWind()
	push := forces.CreatePush(315, 0.2, 30)
	friction := forces.CreateKineticFriction(mover)

	// Main loop
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

		// Update
		mover.ApplyForce(gravity.GetForce())
		mover.ApplyForce(wind.GetForce())
		mover.ApplyForce(friction.GetForce())
		mover.ApplyForce(push.GetForce())

		mover.UpdateSdl2()
		mover.BounceOff()

		// Draw rect
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.DrawRect(rect)

		renderer.Present()
		sdl.Delay(1000 / FrameRate)
	}

	return 0
}
