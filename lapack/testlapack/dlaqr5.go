package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dlaqr5er interface {
	Dlaqr5(wantt, wantz bool, kacc22 int, n, ktop, kbot, nshfts int, sr, si []float64, h []float64, ldh int, iloz, ihiz int, z []float64, ldz int, v []float64, ldv int, u []float64, ldu int, nh int, wh []float64, ldwh int, nv int, wv []float64, ldwv int)
}

func Dlaqr5Test(t *testing.T, impl Dlaqr5er) { _ = "STUB: not implemented"; return }

func testDlaqr5(t *testing.T, impl Dlaqr5er, n, extra, kacc22 int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
