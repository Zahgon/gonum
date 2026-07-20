package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztrmmer interface {
	Ztrmm(side blas.Side, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int)
}

func ZtrmmTest(t *testing.T, impl Ztrmmer) { _ = "STUB: not implemented"; return }

func ztrmmTest(t *testing.T, impl Ztrmmer, side blas.Side, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, m, n int) {
	_ = "STUB: not implemented"
	return
}
