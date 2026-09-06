package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlangter interface {
	Dlangt(norm lapack.MatrixNorm, n int, dl, d, du []float64) float64
}

func DlangtTest(t *testing.T, impl Dlangter) { _ = "STUB: not implemented"; return }

func dlangtTest(t *testing.T, impl Dlangter, rnd *rand.Rand, norm lapack.MatrixNorm, n int) {
	_ = "STUB: not implemented"
	return
}
