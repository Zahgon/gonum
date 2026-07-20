package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtbmver interface {
	Dtbmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float64, lda int, x []float64, incX int)
}

func DtbmvTest(t *testing.T, blasser Dtbmver) { _ = "STUB: not implemented"; return }
