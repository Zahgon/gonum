package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zhpmver interface {
	Zhpmv(uplo blas.Uplo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int)
}

func ZhpmvTest(t *testing.T, impl Zhpmver) { _ = "STUB: not implemented"; return }
