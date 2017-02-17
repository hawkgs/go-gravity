package utils

import (
	"math"
	"testing"
)

func TestRadToDeg(t *testing.T) {
	converted := RadToDeg(math.Pi)

	if converted != 180 {
		t.Error("Expected converted = 180, got", converted)
	}
}

func TestDegToRad(t *testing.T) {
	converted := DegToRad(180)

	if converted != math.Pi {
		t.Error("Expected converted = Pi, got", converted)
	}
}

func TestRoundInLowerHalf(t *testing.T) {
	rounded := Round(0.25)

	if rounded != 0 {
		t.Error("Expected rounded = 0, got", rounded)
	}
}

func TestRoundInHalf(t *testing.T) {
	rounded := Round(0.5)

	if rounded != 1 {
		t.Error("Expected rounded = 1, got", rounded)
	}
}

func TestRoundInUpperHalf(t *testing.T) {
	rounded := Round(0.75)

	if rounded != 1 {
		t.Error("Expected rounded = 1, got", rounded)
	}
}
