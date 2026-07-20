package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zgemmer interface {
	Zgemm(tA, tB blas.Transpose, m, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int)
}

func ZgemmTest(t *testing.T, impl Zgemmer) { _ = "STUB: not implemented"; return }

func zgemmTest(t *testing.T, impl Zgemmer, tA, tB blas.Transpose, m, n, k int) {
	_ = "STUB: not implemented"
	return
}
