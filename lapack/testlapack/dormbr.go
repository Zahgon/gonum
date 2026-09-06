package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dormbrer interface {
	Dormbr(vect lapack.ApplyOrtho, side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int)
	Dgebrder
}

func DormbrTest(t *testing.T, impl Dormbrer) { _ = "STUB: not implemented"; return }
