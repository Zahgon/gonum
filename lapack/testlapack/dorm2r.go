package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dorm2rer interface {
	Dgeqrfer
	Dorm2r(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64)
}

func Dorm2rTest(t *testing.T, impl Dorm2rer) { _ = "STUB: not implemented"; return }
