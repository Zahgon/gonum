package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgetc2er interface {
	Dgetc2(n int, a []float64, lda int, ipiv, jpiv []int) (k int)
}

func Dgetc2Test(t *testing.T, impl Dgetc2er) { _ = "STUB: not implemented"; return }

func dgetc2Test(t *testing.T, impl Dgetc2er, rnd *rand.Rand, n, lda int, perturb bool) {
	_ = "STUB: not implemented"
	return
}
