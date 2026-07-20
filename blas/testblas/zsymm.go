package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zsymmer interface {
	Zsymm(side blas.Side, uplo blas.Uplo, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
}

func ZsymmTest(t *testing.T, impl Zsymmer) { _ = "STUB: not implemented"; return }

func zsymmTest(t *testing.T, impl Zsymmer, side blas.Side, uplo blas.Uplo, m, n int) {
	_ = "STUB: not implemented"
	return
}
