package cmplxNN

import (
	"math"
	"math/cmplx"
)

// Normalize is a function for normalizing the output values of Nueral Network Layers.
// is a variable that can be replaced with another function to allow to alternate Normalization functions
var Normalize = func(c complex128) complex128 {
	abs := cmplx.Abs(c)
	return c * complex(1/abs, 0) * complex(-math.Expm1(-abs), 0)
}
