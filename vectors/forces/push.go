package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// Push force
type Push struct {
	Force
	frames   int
	duration int
}

// CreatePush creates a push force by provided direction (angle in degrees),
// magnitude and duration (in frames)
func CreatePush(direction, magnitude float64, duration int) *Push {
	// todo build the force vector with the direction and magnitude
	return &Push{Force{}, 0, duration}
}

// GetForce returns the force vector of push
func (p *Push) GetForce() *vectors.Vector {
	if p.frames < p.duration {
		p.frames++
		return p.vector
	}
	return nil
}
