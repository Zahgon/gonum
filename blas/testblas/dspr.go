package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsprer interface {
	Dspr(ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64)
}

func DsprTest(t *testing.T, blasser Dsprer) { _ = "STUB: not implemented"; return }
