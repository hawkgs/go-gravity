package forces

import "github.com/hawkgs/go-gravity/vectors/vectors"

// VForce represents the methods of a vector force
type VForce interface {
	GetVector() *vectors.Vector
}

// Force represents the properties of a vector force
type Force struct {
	vector *vectors.Vector
}
