package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrmmer interface {
	Dtrmm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m, n int, alpha float64, a []float64, lda int, b []float64, ldb int)
}

func DtrmmTest(t *testing.T, blasser Dtrmmer) { _ = "STUB: not implemented"; return }
