package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dtrconer interface {
	Dtrcon(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int, work []float64, iwork []int) float64

	Dtrtri(uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int) bool
	Dlantr(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, m, n int, a []float64, lda int, work []float64) float64
}

func DtrconTest(t *testing.T, impl Dtrconer) { _ = "STUB: not implemented"; return }

func dtrconTest(t *testing.T, impl Dtrconer, rnd *rand.Rand, uplo blas.Uplo, diag blas.Diag, n, lda, mattype int) {
	_ = "STUB: not implemented"
	return
}
