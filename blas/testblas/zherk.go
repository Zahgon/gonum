package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zherker interface {
	Zherk(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128, ldc int)
}

func ZherkTest(t *testing.T, impl Zherker) { _ = "STUB: not implemented"; return }

func zherkTest(t *testing.T, impl Zherker, uplo blas.Uplo, trans blas.Transpose, n, k int) {
	_ = "STUB: not implemented"
	return
}
