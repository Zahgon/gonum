package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsymver interface {
	Dsymv(ul blas.Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
}

func DsymvTest(t *testing.T, blasser Dsymver) { _ = "STUB: not implemented"; return }
