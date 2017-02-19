package vectors

import "github.com/veandco/go-sdl2/sdl"

// UpdateSdl2 performs Mover.Update() on SDL2-specific objects
func (m *Mover) UpdateSdl2() {
	switch m.Obj.(type) {
	// SDL_Rect
	case *sdl.Rect:
		rect := m.Obj.(*sdl.Rect)

		m.Update()
		x, y := m.PixelLoc()
		rect.X = int32(x)
		rect.Y = int32(y)
	}
}
