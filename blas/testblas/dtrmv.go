package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrmver interface {
	Dtrmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int)
}

func DtrmvTest(t *testing.T, blasser Dtrmver) { _ = "STUB: not implemented"; return }
