package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlarfer interface {
	Dlarf(side blas.Side, m, n int, v []float64, incv int, tau float64, c []float64, ldc int, work []float64)
}

func DlarfTest(t *testing.T, impl Dlarfer) { _ = "STUB: not implemented"; return }

func runDlarfTest(t *testing.T, impl Dlarfer, side blas.Side) { _ = "STUB: not implemented"; return }

func dlarfTest(t *testing.T, impl Dlarfer, rnd *rand.Rand, side blas.Side, m, n, incv, ldc, nnzv, nnzc int, tau float64) {
	_ = "STUB: not implemented"
	return
}
