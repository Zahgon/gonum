package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsyrer interface {
	Dsyr(ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int)
}

func DsyrTest(t *testing.T, blasser Dsyrer) { _ = "STUB: not implemented"; return }
