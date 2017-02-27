package vectors

import "testing"
import "math"

func TestNewVector(t *testing.T) {
	v := NewVector(1, 2)

	if v.X != 1 || v.Y != 2 {
		t.Error("Expected X = 4 and Y = 2, got", v)
	}
}

func TestAddition(t *testing.T) {
	u := NewVector(1, 2)
	v := NewVector(3, 4)

	u.Add(v)

	if u.X != 4 || u.Y != 6 {
		t.Error("Expected X = 4 and Y = 6, got", u)
	}
}

func TestAdditionWithNegative(t *testing.T) {
	u := NewVector(0, 0)
	v := NewVector(-1, -1)

	u.Add(v)

	if u.X != -1 || u.Y != -1 {
		t.Error("Expected X = -1 and Y = -1, got", u)
	}
}

func TestSubtraction(t *testing.T) {
	u := NewVector(1, 1)
	v := NewVector(-1, -1)

	u.Subtract(v)

	if u.X != 2 || u.Y != 2 {
		t.Error("Expected X = 2 and Y = 2, got", u)
	}
}

func TestMultiplication(t *testing.T) {
	u := NewVector(3, 5)

	u.Multiply(2.5)

	if u.X != 7.5 || u.Y != 12.5 {
		t.Error("Expected X = 7.5 and Y = 12.5, got", u)
	}
}

func TestMultiplicationWithZero(t *testing.T) {
	v := NewVector(7, 9)

	v.Multiply(0)

	if v.X != 0 || v.Y != 0 {
		t.Error("Expected X = 0 and Y = 0, got", v)
	}
}

func TestDivision(t *testing.T) {
	v := NewVector(5, 10)

	v.Divide(2)

	if v.X != 2.5 || v.Y != 5 {
		t.Error("Expected X = 2.5 and Y = 5, got", v)
	}
}

func TestMagnitude(t *testing.T) {
	v := NewVector(3, 4)
	mag := v.Magnitude()

	if mag != 5 {
		t.Error("Expected ||v|| = 5, got", mag)
	}
}

func TestLimitWithHigherMag(t *testing.T) {
	v := NewVector(3, 4)
	v.Limit(3)
	mag := v.Magnitude()

	var ratio, x, y float64
	ratio = 5.0 / 3.0
	x = 3 / ratio
	y = 4 / ratio

	if mag != 3 {
		t.Error("Expected ||v|| = 3, got", mag)
	}

	if v.X != x || v.Y != y {
		t.Error("Expected X = 9/5 and Y = 12/5, got", v)
	}
}

func TestLimitWithLowerMag(t *testing.T) {
	v := NewVector(3, 4)
	v.Limit(10)
	mag := v.Magnitude()

	if mag != 5 {
		t.Error("Expected ||v|| = 5, got", mag)
	}

	if v.X != 3 || v.Y != 4 {
		t.Error("Expected X = 3 and Y = 4, got", v)
	}
}

func TestDistance(t *testing.T) {
	u := NewVector(1, 1)
	v := NewVector(4, 4)

	distance := u.Distance(v)
	expected := 3 * math.Sqrt(2)

	if math.Abs(distance-expected) > 0.000000000000001 {
		t.Error("Expected distance = ", expected, ", got", distance)
	}
}
