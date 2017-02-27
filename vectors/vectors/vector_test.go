package vectors

import "testing"
import "math"
import "fmt"

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

func TestHeading(t *testing.T) {
	tests := []struct {
		input  *Vector
		result float64
	}{
		{input: NewVector(1, 0), result: 0},
		{input: NewVector(1, 1), result: 45},
		{input: NewVector(0, 1), result: 90},
		{input: NewVector(-1, 1), result: 135},
		{input: NewVector(-1, 0), result: 180},
		{input: NewVector(-1, -1), result: 225},
		{input: NewVector(0, -1), result: 270},
		{input: NewVector(1, -1), result: 315},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Heading %v", test.result), func(t *testing.T) {
			heading := test.input.Heading()

			if heading != test.result {
				t.Error("Expected heading to be", test.result, ", got", heading)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	vector := NewVector(1, 0)
	tests := []struct {
		rotate float64
		expect *Vector
	}{
		{rotate: 0, expect: NewVector(1, 0)},
		{rotate: 45, expect: NewVector(1, 1)},
		{rotate: 90, expect: NewVector(0, 1)},
		{rotate: 135, expect: NewVector(-1, 1)},
		{rotate: 180, expect: NewVector(-1, 0)},
		{rotate: 225, expect: NewVector(-1, -1)},
		{rotate: 270, expect: NewVector(0, -1)},
		{rotate: 315, expect: NewVector(1, -1)},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Rotate with %v", test.rotate), func(t *testing.T) {
			v := vector.Copy()
			v.Rotate(test.rotate)

			if v.X != test.expect.X || v.Y != test.expect.Y {
				t.Error("Expected vector to be", test.expect, ", got", v)
			}
		})
	}
}

func TestNormalize(t *testing.T) {
	vector := NewVector(5, 5)
	vector.Normalize()
	mag := vector.Magnitude()

	if !(0.9999999999999 < mag && mag <= 1) {
		t.Error("Expected magnitude to be 1, got", mag)
	}

	side := 1 / math.Sqrt(2)

	if vector.X != side || vector.Y != side {
		t.Error("Expected X = Y =", side, ", got", vector)
	}
}

func TestDistance(t *testing.T) {
	u := NewVector(1, 1)
	v := NewVector(4, 4)

	distance := u.Distance(v)
	expected := 3 * math.Sqrt(2)

	if math.Abs(distance-expected) > 0.000000000000001 {
		t.Error("Expected distance to be", expected, ", got", distance)
	}
}

func TestCopy(t *testing.T) {
	u := NewVector(1, 2)
	v := u.Copy()

	if u.X != v.X || u.Y != v.Y {
		t.Error("Expected X = 1 and Y = 2, got", v)
	}

	if u == v {
		t.Error("Expected U and V to have different addresses, got", &u, "and", &v)
	}
}
