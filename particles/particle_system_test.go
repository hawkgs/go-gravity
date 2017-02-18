package particles

import (
	"testing"

	"github.com/hAWKdv/go-gravity/vectors/forces"
	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

func TestNewConf(t *testing.T) {
	loc := vectors.NewVector(0, 0)
	cont := vectors.NewVector(10, 10)

	conf := NewConf(true, 10, 20, loc, cont)

	if conf.continious != true ||
		conf.particleLifespan != 10 ||
		conf.maxRetardation != 20 ||
		conf.location != loc ||
		conf.container != cont {
		t.Error("Expected &{true 10 20 &loc &cont}, got", conf)
	}
}

func TestNewParticleSystem(t *testing.T) {
	strings := []string{"a", "b"}
	vector := vectors.NewVector(0, 0)
	conf := NewConf(true, 1, 1, vector, vector)

	ps := NewParticleSystem(strings, conf)

	if ps.particles[0].mover.Obj != "a" || ps.particles[1].mover.Obj != "b" {
		t.Error("Particle mover objects do not match A and B")
	}

	if ps.conf != conf {
		t.Error("Expected conf to match")
	}

	if ps.frame != 0 {
		t.Error("Expected frame to be 0, got", ps.frame)
	}

	gravity := forces.CreateGravity().GetForce()
	psGravity := ps.gravity.GetForce()

	if gravity.X != psGravity.X || gravity.Y != psGravity.Y {
		t.Error("Expected gravity values to match", gravity, ", got", psGravity)
	}
}

func TestIncorrectSetMassesInput(t *testing.T) {
	strings := []string{"a", "b"}
	vector := vectors.NewVector(0, 0)
	conf := NewConf(true, 1, 1, vector, vector)

	ps := NewParticleSystem(strings, conf)
	err := ps.SetMasses([]float64{2, 3, 4, 5})

	if err == nil {
		t.Error("Expected an error from input slice")
	}
}
