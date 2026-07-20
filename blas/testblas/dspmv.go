package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dspmver interface {
	Dspmv(ul blas.Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int)
}

func DspmvTest(t *testing.T, blasser Dspmver) { _ = "STUB: not implemented"; return }
