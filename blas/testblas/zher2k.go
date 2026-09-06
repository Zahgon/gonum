package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zher2ker interface {
	Zher2k(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta float64, c []complex128, ldc int)
}

func Zher2kTest(t *testing.T, impl Zher2ker) { _ = "STUB: not implemented"; return }

func zher2kTest(t *testing.T, impl Zher2ker, uplo blas.Uplo, trans blas.Transpose, n, k int) {
	_ = "STUB: not implemented"
	return
}
