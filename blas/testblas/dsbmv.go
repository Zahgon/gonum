package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsbmver interface {
	Dsbmv(ul blas.Uplo, n, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
}

func DsbmvTest(t *testing.T, blasser Dsbmver) { _ = "STUB: not implemented"; return }
