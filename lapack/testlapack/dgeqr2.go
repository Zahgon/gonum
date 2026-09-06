package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgeqr2er interface {
	Dgeqr2(m, n int, a []float64, lda int, tau []float64, work []float64)
}

func Dgeqr2Test(t *testing.T, impl Dgeqr2er) { _ = "STUB: not implemented"; return }

func dgeqr2Test(t *testing.T, impl Dgeqr2er, rnd *rand.Rand, m, n, lda int) {
	_ = "STUB: not implemented"
	return
}
