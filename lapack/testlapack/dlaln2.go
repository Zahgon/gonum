package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dlaln2er interface {
	Dlaln2(trans bool, na, nw int, smin, ca float64, a []float64, lda int, d1, d2 float64, b []float64, ldb int, wr, wi float64, x []float64, ldx int) (scale, xnorm float64, ok bool)
}

func Dlaln2Test(t *testing.T, impl Dlaln2er) { _ = "STUB: not implemented"; return }

func testDlaln2(t *testing.T, impl Dlaln2er, trans bool, na, nw, extra int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
