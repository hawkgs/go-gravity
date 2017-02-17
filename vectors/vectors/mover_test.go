package vectors

import "testing"

func TestNewMover(t *testing.T) {
	vector := NewVector(1, 1)
	m := NewMover("a", vector, vector)

	if m.Obj != "a" ||
		m.acceleration.X != 0 || m.acceleration.Y != 0 ||
		m.velocity.X != 0 || m.velocity.Y != 0 ||
		m.location.X != 1 || m.location.Y != 1 ||
		m.container.X != 1 || m.container.Y != 1 ||
		m.initLoc.X != 1 || m.initLoc.Y != 1 ||
		m.mass != 1 || m.limit != 0 {
		t.Error("Incorrect Mover configuration", m)
	}
}

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

func TestApplyForce(t *testing.T) {
	location := NewVector(0, 0)
	cont := NewVector(1, 1)
	force1 := NewVector(0.5, 0)
	force2 := NewVector(0, 0.5)
	mover := NewMover(nil, location, cont)

	mover.ApplyForce(force1)
	mover.ApplyForce(force2)

	if mover.acceleration.X != 0.5 || mover.acceleration.Y != 0.5 {
		t.Error("Expected acceleration X and Y to be 0.5, got", mover.acceleration)
	}
}

func TestGetVelocityAndUpdate(t *testing.T) {
	location := NewVector(0, 0)
	cont := NewVector(1, 1)
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

func TestPixelLoc(t *testing.T) {
	vector := NewVector(2.4, 5.7)
	mover := NewMover(nil, vector, vector)

	x, y := mover.PixelLoc()

	if x != 2 || y != 6 {
		t.Error("Expected X and Y, to be 2 and 6, got", x, y)
	}
}

// Graphical representation
// ______
// ../\..
func TestBounceOffTop(t *testing.T) {
	location := NewVector(5, 0)
	cont := NewVector(10, 10)
	force := NewVector(1, -1)
	mover := NewMover(nil, location, cont)

	mover.ApplyForce(force)
	mover.Update()
	mover.BounceOff()

	if mover.location.X != 6 || mover.location.Y != 0 {
		t.Error("Expected location to be {6, 0}, got", mover.location)
	}

	if mover.velocity.X != 1 || mover.velocity.Y != 1 {
		t.Error("Expected velocity to be {1, 1}, got", mover.velocity)
	}
}

// Graphical representation
// ..\/..
// ¯¯¯¯¯¯
func TestBounceOffBottom(t *testing.T) {
	location := NewVector(5, 10)
	cont := NewVector(10, 10)
	force := NewVector(1, 1)
	mover := NewMover(nil, location, cont)

	mover.ApplyForce(force)
	mover.Update()
	mover.BounceOff()

	if mover.location.X != 6 || mover.location.Y != 10 {
		t.Error("Expected location to be {6, 10}, got", mover.location)
	}

	if mover.velocity.X != 1 || mover.velocity.Y != -1 {
		t.Error("Expected velocity to be {1, -1}, got", mover.velocity)
	}
}

// Graphical representation
// |<
func TestBounceOffLeft(t *testing.T) {
	location := NewVector(0, 5)
	cont := NewVector(10, 10)
	force := NewVector(-1, 1)
	mover := NewMover(nil, location, cont)

	mover.ApplyForce(force)
	mover.Update()
	mover.BounceOff()

	if mover.location.X != 0 || mover.location.Y != 6 {
		t.Error("Expected location to be {0, 6}, got", mover.location)
	}

	if mover.velocity.X != 1 || mover.velocity.Y != 1 {
		t.Error("Expected velocity to be {1, 1}, got", mover.velocity)
	}
}

// Graphical representation
// >|
func TestBounceOffRight(t *testing.T) {
	location := NewVector(10, 5)
	cont := NewVector(10, 10)
	force := NewVector(1, 1)
	mover := NewMover(nil, location, cont)

	mover.ApplyForce(force)
	mover.Update()
	mover.BounceOff()

	if mover.location.X != 10 || mover.location.Y != 6 {
		t.Error("Expected location to be {10, 6}, got", mover.location)
	}

	if mover.velocity.X != -1 || mover.velocity.Y != 1 {
		t.Error("Expected velocity to be {-1, 1}, got", mover.velocity)
	}
}

func TestReset(t *testing.T) {
	location := NewVector(5, 5)
	force := NewVector(3, 3)
	mover := NewMover(nil, location, location)

	mover.ApplyForce(force)
	mover.Update()
	mover.Reset()

	if mover.location.X != 5 || mover.location.Y != 5 {
		t.Error("Expected location to be {5, 5}, got", mover.location)
	}

	if mover.velocity.X != 0 || mover.velocity.Y != 0 {
		t.Error("Expected location to be {0, 0}, got", mover.velocity)
	}
}
