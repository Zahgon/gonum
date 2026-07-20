package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgehd2er interface {
	Dgehd2(n, ilo, ihi int, a []float64, lda int, tau, work []float64)
}

func Dgehd2Test(t *testing.T, impl Dgehd2er) { _ = "STUB: not implemented"; return }

func testDgehd2(t *testing.T, impl Dgehd2er, n, extra int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
