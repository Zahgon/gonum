package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrsmer interface {
	Dtrsm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m, n int,
		alpha float64, a []float64, lda int, b []float64, ldb int)
}

func DtrsmTest(t *testing.T, impl Dtrsmer) { _ = "STUB: not implemented"; return }
