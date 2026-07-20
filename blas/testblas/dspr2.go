package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dspr2er interface {
	Dspr2(ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64)
}

func Dspr2Test(t *testing.T, blasser Dspr2er) { _ = "STUB: not implemented"; return }
