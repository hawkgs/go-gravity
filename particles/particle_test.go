package particles

import (
	"testing"

	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

func TestNewParticle(t *testing.T) {
	vector := vectors.NewVector(0, 0)
	mover := vectors.NewMover(nil, vector, vector)

	particle := NewParticle(mover, 100, 100)

	if particle.retardation > 100 {
		t.Error("The retardation overflows the max")
	}

	if particle.mover != mover ||
		particle.push == nil ||
		particle.friction == nil ||
		particle.lifespan != 100+particle.retardation ||
		particle.id == 0 {
		t.Error("Incorrect particle", particle)
	}

}
