package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dormlqer interface {
	Dorml2er
	Dormlq(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int)
}

func DormlqTest(t *testing.T, impl Dormlqer) { _ = "STUB: not implemented"; return }
