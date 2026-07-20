package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dgbmver interface {
	Dgbmv(tA blas.Transpose, m, n, kL, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int)
}

func DgbmvTest(t *testing.T, blasser Dgbmver) { _ = "STUB: not implemented"; return }
