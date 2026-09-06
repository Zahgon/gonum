package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztrsmer interface {
	Ztrsm(side blas.Side, uplo blas.Uplo, transA blas.Transpose, diag blas.Diag, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int)
}

func ZtrsmTest(t *testing.T, impl Ztrsmer) { _ = "STUB: not implemented"; return }

func ztrsmTest(t *testing.T, impl Ztrsmer, side blas.Side, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, m, n int) {
	_ = "STUB: not implemented"
	return
}
