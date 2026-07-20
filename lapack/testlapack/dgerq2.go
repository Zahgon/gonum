package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgerq2er interface {
	Dgerq2(m, n int, a []float64, lda int, tau []float64, work []float64)
}

func Dgerq2Test(t *testing.T, impl Dgerq2er) { _ = "STUB: not implemented"; return }

func dgerq2Test(t *testing.T, impl Dgerq2er, rnd *rand.Rand, m, n, lda int) {
	_ = "STUB: not implemented"
	return
}
