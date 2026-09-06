package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zhpr2er interface {
	Zhpr2(uplo blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128)
}

func Zhpr2Test(t *testing.T, impl Zhpr2er) { _ = "STUB: not implemented"; return }
