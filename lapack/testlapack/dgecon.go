package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dgeconer interface {
	Dgecon(norm lapack.MatrixNorm, n int, a []float64, lda int, anorm float64, work []float64, iwork []int) float64

	Dgetrier
	Dlanger
}

func DgeconTest(t *testing.T, impl Dgeconer) { _ = "STUB: not implemented"; return }

func dgeconTest(t *testing.T, impl Dgeconer, rnd *rand.Rand, n, lda int) {
	_ = "STUB: not implemented"
	return
}
