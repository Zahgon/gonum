package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztpmver interface {
	Ztpmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, ap []complex128, x []complex128, incX int)
}

func ZtpmvTest(t *testing.T, impl Ztpmver) { _ = "STUB: not implemented"; return }
