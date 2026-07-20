package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlangber interface {
	Dlangb(norm lapack.MatrixNorm, m, n, kl, ku int, ab []float64, ldab int) float64
}

func DlangbTest(t *testing.T, impl Dlangber) { _ = "STUB: not implemented"; return }

func dlangbTest(t *testing.T, impl Dlangber, rnd *rand.Rand, norm lapack.MatrixNorm, m, n, kl, ku, ldab int) {
	_ = "STUB: not implemented"
	return
}
