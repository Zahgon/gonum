package testblas

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztrsver interface {
	Ztrsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, a []complex128, lda int, x []complex128, incX int)

	Ztrmver
}

func ZtrsvTest(t *testing.T, impl Ztrsver) { _ = "STUB: not implemented"; return }

func ztrsvTest(t *testing.T, impl Ztrsver, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, lda, incX int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
