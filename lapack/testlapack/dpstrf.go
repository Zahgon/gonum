package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpstrfer interface {
	Dpstrf(uplo blas.Uplo, n int, a []float64, lda int, piv []int, tol float64, work []float64) (rank int, ok bool)
}

func DpstrfTest(t *testing.T, impl Dpstrfer) { _ = "STUB: not implemented"; return }

func dpstrfTest(t *testing.T, impl Dpstrfer, rnd *rand.Rand, uplo blas.Uplo, n, lda, rankWant int) {
	_ = "STUB: not implemented"
	return
}

func residualDpstrf(uplo blas.Uplo, n int, a, aFac []float64, lda int, rank int, piv []int) float64 {
	_ = "STUB: not implemented"
	return 0
}
