package vectors

import (
	"testing"

	"github.com/hAWKdv/go-gravity/vectors/vectors"
)

func TestAddition(t *testing.T) {
	u := vectors.NewVector(1, 2)
	v := vectors.NewVector(3, 4)

	u.Add(v)

	if u.X != 4 || u.Y != 6 {
		t.Error("Expected X = 4 and Y = 6, got", u)
	}
}
