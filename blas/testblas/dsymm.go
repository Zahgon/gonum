package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsymmer interface {
	Dsymm(s blas.Side, ul blas.Uplo, m, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int)
}

func DsymmTest(t *testing.T, blasser Dsymmer) { _ = "STUB: not implemented"; return }
