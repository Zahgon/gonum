package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlantrer interface {
	Dlanger
	Dlantr(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, m, n int, a []float64, lda int, work []float64) float64
}

func DlantrTest(t *testing.T, impl Dlantrer) { _ = "STUB: not implemented"; return }

func dlantrTest(t *testing.T, impl Dlantrer, rnd *rand.Rand, uplo blas.Uplo, diag blas.Diag, m, n, lda int) {
	_ = "STUB: not implemented"
	return
}
