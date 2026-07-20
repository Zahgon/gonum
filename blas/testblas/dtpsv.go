package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtpsver interface {
	Dtpsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int)
}

func DtpsvTest(t *testing.T, blasser Dtpsver) { _ = "STUB: not implemented"; return }
