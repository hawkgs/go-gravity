package forces

import (
	"math"

	"github.com/hAWKdv/go-gravity/vectors/utils"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

// Push force
type Push struct {
	Force
	frames   int
	duration int
}

// CreatePush creates a push force by provided direction (angle in degrees),
// magnitude and duration (in frames)
func CreatePush(direction, magnitude float64, duration int) *Push {
	rad := utils.DegToRad(direction)
	vector := vectors.NewVector(magnitude*math.Cos(rad), magnitude*math.Sin(rad))
	return &Push{Force{vector}, 0, duration}
}

// GetForce returns the force vector of push
func (p *Push) GetForce() *vectors.Vector {
	if p.frames < p.duration {
		p.frames++
		return p.vector
	}
	return nil
}
