package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlantber interface {
	Dlantb(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, n, k int, a []float64, lda int, work []float64) float64
}

func DlantbTest(t *testing.T, impl Dlantber) { _ = "STUB: not implemented"; return }

func dlantbTest(t *testing.T, impl Dlantber, rnd *rand.Rand, norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, n, k, lda int) {
	_ = "STUB: not implemented"
	return
}
