package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsyr2ker interface {
	Dsyr2k(ul blas.Uplo, tA blas.Transpose, n, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
}

func Dsyr2kTest(t *testing.T, blasser Dsyr2ker) { _ = "STUB: not implemented"; return }
