package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zgemver interface {
	Zgemv(trans blas.Transpose, m, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int)
}

func ZgemvTest(t *testing.T, impl Zgemver) { _ = "STUB: not implemented"; return }
