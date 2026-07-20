package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zsyrker interface {
	Zsyrk(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex128, a []complex128, lda int, beta complex128, c []complex128, ldc int)
}

func ZsyrkTest(t *testing.T, impl Zsyrker) { _ = "STUB: not implemented"; return }

func zsyrkTest(t *testing.T, impl Zsyrker, uplo blas.Uplo, trans blas.Transpose, n, k int) {
	_ = "STUB: not implemented"
	return
}
