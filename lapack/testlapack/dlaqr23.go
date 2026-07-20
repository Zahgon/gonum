package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

type Dlaqr23er interface {
	Dlaqr23(wantt, wantz bool, n, ktop, kbot, nw int, h []float64, ldh int, iloz, ihiz int, z []float64, ldz int, sr, si []float64, v []float64, ldv int, nh int, t []float64, ldt int, nv int, wv []float64, ldwv int, work []float64, lwork int, recur int) (ns, nd int)
}

type dlaqr23Test struct {
	wantt, wantz bool
	ktop, kbot   int
	nw           int
	h            blas64.General
	iloz, ihiz   int

	evWant []complex128
}

func newDlaqr23TestCase(wantt, wantz bool, n, ldh int, rnd *rand.Rand) dlaqr23Test {
	_ = "STUB: not implemented"
	return *new(dlaqr23Test)
}

func Dlaqr23Test(t *testing.T, impl Dlaqr23er) { _ = "STUB: not implemented"; return }

func testDlaqr23(t *testing.T, impl Dlaqr23er, test dlaqr23Test, opt bool, recur int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
