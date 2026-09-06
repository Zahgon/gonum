package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dormr2er interface {
	Dgerqf(m, n int, a []float64, lda int, tau, work []float64, lwork int)
	Dormr2(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64)
}

func Dormr2Test(t *testing.T, impl Dormr2er) { _ = "STUB: not implemented"; return }
