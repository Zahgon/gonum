package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dormqrer interface {
	Dorm2rer
	Dormqr(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int)
}

func DormqrTest(t *testing.T, impl Dormqrer) { _ = "STUB: not implemented"; return }
