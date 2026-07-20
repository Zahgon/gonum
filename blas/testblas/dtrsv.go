package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrsver interface {
	Dtrsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int)
}

func DtrsvTest(t *testing.T, blasser Dtrsver) { _ = "STUB: not implemented"; return }
