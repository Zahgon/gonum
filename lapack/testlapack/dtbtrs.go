package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtbtrser interface {
	Dtbtrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, kd, nrhs int, a []float64, lda int, b []float64, ldb int) bool
}

func DtbtrsTest(t *testing.T, impl Dtbtrser) { _ = "STUB: not implemented"; return }

func dtbtrsTest(t *testing.T, impl Dtbtrser, rnd *rand.Rand, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, kd, nrhs int, lda, ldb int, singular bool) {
	_ = "STUB: not implemented"
	return
}
