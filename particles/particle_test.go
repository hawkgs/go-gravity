package particles

import (
	"testing"

	"github.com/hawkgs/go-gravity/vectors/vectors"
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

func TestParticleReset(t *testing.T) {
	vector := vectors.NewVector(0, 0)
	mover := vectors.NewMover(nil, vector, vector)

	particle := NewParticle(mover, 100, 100)
	particle.mover.ApplyForce(vectors.NewVector(1, 2))
	particle.mover.Update()

	particle.Reset(20)

	if particle.lifespan != 20 {
		t.Error("Expected lifespan to be 20, got", 20)
	}

	if particle.mover.GetVelocity().X != 0 || particle.mover.GetVelocity().Y != 0 {
		t.Error("Expected mover's velocity to be {0 0}, got", particle.mover)
	}
}
