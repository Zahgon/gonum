package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zhemmer interface {
	Zhemm(side blas.Side, uplo blas.Uplo, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
}

func ZhemmTest(t *testing.T, impl Zhemmer) { _ = "STUB: not implemented"; return }

func zhemmTest(t *testing.T, impl Zhemmer, side blas.Side, uplo blas.Uplo, m, n int) {
	_ = "STUB: not implemented"
	return
}
