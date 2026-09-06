package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrtrser interface {
	Dtrtrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, nrhs int, a []float64, lda int, b []float64, ldb int) bool
}

func DtrtrsTest(t *testing.T, impl Dtrtrser) { _ = "STUB: not implemented"; return }

func dtrtrsTest(t *testing.T, impl Dtrtrser, rnd *rand.Rand, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, nrhs int, lda, ldb int, singular bool) {
	_ = "STUB: not implemented"
	return
}
