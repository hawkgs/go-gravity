package vectors

import "testing"

func TestSetMass(t *testing.T) {
	vector := NewVector(0, 0)
	mover := NewMover(nil, vector, vector)
	mover.SetMass(20)

	if mover.mass != 20 {
		t.Error("Expected mass to be 20, got", mover.mass)
	}
}

func TestSmallerThanOneMass(t *testing.T) {
	vector := NewVector(0, 0)
	mover := NewMover(nil, vector, vector)
	mover.SetMass(0.5)

	if mover.mass != 1 {
		t.Error("Expected mass to be 1, got", mover.mass)
	}
}

func TestSetLimit(t *testing.T) {
	location := NewVector(0, 0)
	cont := NewVector(1000, 1)
	force := NewVector(0.1, 0)
	mover := NewMover(nil, location, cont)
	mover.SetLimit(3)

	for i := 0; i < 100; i++ {
		mover.ApplyForce(force)
		mover.Update()
	}

	v := mover.GetVelocity().Magnitude()
	if v != 3 {
		t.Error("Expected velocity to be 3, got", v)
	}
}

// func TestApplyForce(t *testing.T) {

// }

func TestGetVelocityAndUpdate(t *testing.T) {
	location := NewVector(0, 0)
	cont := NewVector(1000, 1)
	force1 := NewVector(0.2, 0)
	force2 := NewVector(0.3, 0)
	mover := NewMover(nil, location, cont)

	mover.ApplyForce(force1)
	mover.ApplyForce(force2)
	mover.Update()

	v := mover.GetVelocity()

	if v.X != 0.5 {
		t.Error("Expected v.X to be 0.5, got", v.X)
	}
}
