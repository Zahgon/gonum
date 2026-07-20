package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsyker interface {
	Dsyrk(ul blas.Uplo, tA blas.Transpose, n, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int)
}

func DsyrkTest(t *testing.T, blasser Dsyker) { _ = "STUB: not implemented"; return }
