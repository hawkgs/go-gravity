package forces

import (
	"math"

	"github.com/hawkgs/go-gravity/vectors/utils"
	"github.com/hawkgs/go-gravity/vectors/vectors"
)

// Wind force
type Wind struct {
	Force
}

// CreateWind creates a wind force
func CreateWind(direction, magnitude float64) *Wind {
	rad := utils.DegToRad(direction)
	vector := vectors.NewVector(magnitude*math.Cos(rad), magnitude*math.Sin(rad))

	return &Wind{Force{vector}}
}

// GetForce returns the force vector of gravity
func (w *Wind) GetForce() *vectors.Vector {
	return w.vector
}
