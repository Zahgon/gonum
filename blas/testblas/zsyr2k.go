package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zsyr2ker interface {
	Zsyr2k(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
}

func Zsyr2kTest(t *testing.T, impl Zsyr2ker) { _ = "STUB: not implemented"; return }

func zsyr2kTest(t *testing.T, impl Zsyr2ker, uplo blas.Uplo, trans blas.Transpose, n, k int) {
	_ = "STUB: not implemented"
	return
}
