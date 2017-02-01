package forces

import "github.com/hAWKdv/go-gravity/vectors/vectors"

// VForce describes
type VForce interface {
	GetVector() vectors.Vector
}

// Force test
type Force struct {
	vector *vectors.Vector
}
