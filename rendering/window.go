package rendering

import (
	"fmt"
	"os"

	vector "github.com/hAWKdv/go-gravity/vectors"
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
	mover := vector.NewMover(&rect, vector.NewVector(float64(rect.X), float64(rect.Y)), vector.NewVector(WindowWidth, WindowHeight))
	// mover.SetLimit(0.2)
	gravity := vector.NewVector(0, 0.08)
	wind := vector.NewVector(0.01, 0)

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
		mover.ApplyForce(gravity)
		mover.ApplyForce(wind)
		mover.UpdateSdl2()
		mover.BounceOff()

		// rect.X = rect.X + 1
		// rect.Y = rect.Y + 1
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.DrawRect(&rect)

		renderer.Present()
		sdl.Delay(1000 / FrameRate)
	}

	return 0
}
