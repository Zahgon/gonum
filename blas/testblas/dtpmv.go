package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtpmver interface {
	Dtpmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int)
}

func DtpmvTest(t *testing.T, blasser Dtpmver) { _ = "STUB: not implemented"; return }
