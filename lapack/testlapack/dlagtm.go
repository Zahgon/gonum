package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlagtmer interface {
	Dlagtm(trans blas.Transpose, m, n int, alpha float64, dl, d, du []float64, b []float64, ldb int, beta float64, c []float64, ldc int)
}

func DlagtmTest(t *testing.T, impl Dlagtmer) { _ = "STUB: not implemented"; return }

func dlagtmTest(t *testing.T, impl Dlagtmer, rnd *rand.Rand, trans blas.Transpose, m, n int, ldb, ldc int, alpha, beta float64) {
	_ = "STUB: not implemented"
	return
}
