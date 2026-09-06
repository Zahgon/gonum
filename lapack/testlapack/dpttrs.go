package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

type Dpttrser interface {
	Dpttrs(n, nrhs int, d, e []float64, b []float64, ldb int)

	Dpttrfer
}

func DpttrsTest(t *testing.T, impl Dpttrser) { _ = "STUB: not implemented"; return }

func dpttrsTest(t *testing.T, impl Dpttrser, rnd *rand.Rand, n, nrhs, ldb int) {
	_ = "STUB: not implemented"
	return
}

func dstmm(m, n int, d, e []float64, b []float64, ldb int, c []float64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func dpttrsResidual(xGot, xWant blas64.General) float64 { _ = "STUB: not implemented"; return 0 }
