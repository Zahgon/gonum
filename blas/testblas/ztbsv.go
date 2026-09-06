package testblas

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Ztbsver interface {
	Ztbsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k int, ab []complex128, ldab int, x []complex128, incX int)

	Ztbmver
}

func ZtbsvTest(t *testing.T, impl Ztbsver) { _ = "STUB: not implemented"; return }

func ztbsvTest(t *testing.T, impl Ztbsver, rnd *rand.Rand, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k, ldab, incX int) {
	_ = "STUB: not implemented"
	return
}
