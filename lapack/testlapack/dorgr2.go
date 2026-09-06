package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dorgr2er interface {
	Dorgr2(m, n, k int, a []float64, lda int, tau []float64, work []float64)

	Dgerqfer
}

func Dorgr2Test(t *testing.T, impl Dorgr2er) { _ = "STUB: not implemented"; return }

func dorgr2Test(t *testing.T, impl Dorgr2er, rnd *rand.Rand, m, n, k, lda int) {
	_ = "STUB: not implemented"
	return
}
