package rendering

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

// Main window variables
const (
	WindowTitle  = "go-gravity"
	WindowWidth  = 800
	WindowHeight = 700
	FrameRate    = 60
)

// CreateSdlWindow creates SDL2 window
func CreateSdlWindow() (*sdl.Window, int) {
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

// CreateSdlRenderer creates SDL2 renderer
func CreateSdlRenderer(window *sdl.Window) (*sdl.Renderer, int) {
	exit := 0
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		exit = 2
	}

	return renderer, exit
}
