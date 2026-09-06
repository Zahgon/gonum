package testblas

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztpsver interface {
	Ztpsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, ap []complex128, x []complex128, incX int)

	Ztpmver
}

func ZtpsvTest(t *testing.T, impl Ztpsver) { _ = "STUB: not implemented"; return }

func ztpsvTest(t *testing.T, impl Ztpsver, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, incX int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
