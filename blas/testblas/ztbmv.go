package testblas

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztbmver interface {
	Ztbmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k int, ab []complex128, ldab int, x []complex128, incX int)

	Ztrmver
}

func ZtbmvTest(t *testing.T, impl Ztbmver) { _ = "STUB: not implemented"; return }

func testZtbmv(t *testing.T, impl Ztbmver, rnd *rand.Rand, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k, ldab, incX int) {
	_ = "STUB: not implemented"
	return
}
