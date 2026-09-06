package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsyr2er interface {
	Dsyr2(ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int)
}

func Dsyr2Test(t *testing.T, blasser Dsyr2er) { _ = "STUB: not implemented"; return }
