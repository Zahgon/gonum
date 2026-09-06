package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dlasy2er interface {
	Dlasy2(tranl, tranr bool, isgn, n1, n2 int, tl []float64, ldtl int, tr []float64, ldtr int, b []float64, ldb int, x []float64, ldx int) (scale, xnorm float64, ok bool)
}

func Dlasy2Test(t *testing.T, impl Dlasy2er) { _ = "STUB: not implemented"; return }

func testDlasy2(t *testing.T, impl Dlasy2er, tranl, tranr bool, isgn, n1, n2, extra int, big bool, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
