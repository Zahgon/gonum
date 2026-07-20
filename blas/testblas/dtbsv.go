package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtbsver interface {
	Dtbsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float64, lda int, x []float64, incX int)
	Dtrsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int)
}

func DtbsvTest(t *testing.T, blasser Dtbsver) { _ = "STUB: not implemented"; return }
