package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zhprer interface {
	Zhpr(uplo blas.Uplo, n int, alpha float64, x []complex128, incX int, ap []complex128)
}

func ZhprTest(t *testing.T, impl Zhprer) { _ = "STUB: not implemented"; return }
