package cmplxNN

import (
	"encoding/json"
	"math"
	"math/cmplx"
)

// Normalizer is an interface to allow alternate functions for normalizing output
// it is recommended that the type that implements Normalizer also implements
// json.Marshaler if you want to save the Nueral Net to a file.
type Normalizer interface {
	Normalize(complex128) complex128
}

type ExponentialNormalizer struct{}

func (en ExponentialNormalizer) Normalize(c complex128) complex128 {
	abs := cmplx.Abs(c)
	return c * complex(1/abs, 0) * complex(-math.Expm1(-abs), 0)
}

func (en ExponentialNormalizer) MarshalJSON() (b []byte, e error) {
	return json.Marshal("ExponentialNormalizer")
}
