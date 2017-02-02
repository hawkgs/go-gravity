package rendering

import (
	"fmt"
	"os"

	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
	"github.com/veandco/go-sdl2/sdl"
)

// Main window variables
const (
	WindowTitle  = "go-gravity"
	WindowWidth  = 800
	WindowHeight = 700
	FrameRate    = 60
)

func createSdlWindow() (*sdl.Window, int) {
	exit := 0
	window, err := sdl.CreateWindow(
		WindowTitle,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WindowWidth,
		WindowHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		exit = 1
	}

	return window, exit
}

func createSdlRenderer(window *sdl.Window) (*sdl.Renderer, int) {
	exit := 0
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		exit = 2
	}

	return renderer, exit
}

// Render runs the main rendering process
func Render() int {
	// Create window
	window, exit := createSdlWindow()
	if exit != 0 {
		return exit
	}
	defer window.Destroy()

	// Create renderer
	renderer, exit := createSdlRenderer(window)
	if exit != 0 {
		return exit
	}
	defer renderer.Destroy()
	renderer.Clear()

	// Create rect
	rect := sdl.Rect{X: 100, Y: 100, W: 10, H: 10}
	mover := vectors.NewMover(
		&rect,
		vectors.NewVector(float64(rect.X), float64(rect.Y)),
		vectors.NewVector(WindowWidth, WindowHeight))
	// mover.SetLimit(0.2)

	// Create some forces
	gravity := forces.CreateGravity()
	// wind := forces.CreateWind()
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
		// mover.ApplyForce(wind.GetForce())
		mover.ApplyForce(friction.GetForce())
		mover.ApplyForce(push.GetForce())
		mover.UpdateSdl2()
		mover.BounceOff()

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.DrawRect(&rect)

		renderer.Present()
		sdl.Delay(1000 / FrameRate)
	}

	return 0
}
