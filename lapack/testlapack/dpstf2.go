package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpstf2er interface {
	Dpstf2(uplo blas.Uplo, n int, a []float64, lda int, piv []int, tol float64, work []float64) (rank int, ok bool)
}

func Dpstf2Test(t *testing.T, impl Dpstf2er) { _ = "STUB: not implemented"; return }

func dpstf2Test(t *testing.T, impl Dpstf2er, rnd *rand.Rand, uplo blas.Uplo, n, lda, rankWant int) {
	_ = "STUB: not implemented"
	return
}
