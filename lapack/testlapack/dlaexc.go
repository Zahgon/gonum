package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dlaexcer interface {
	Dlaexc(wantq bool, n int, t []float64, ldt int, q []float64, ldq int, j1, n1, n2 int, work []float64) bool
}

func DlaexcTest(t *testing.T, impl Dlaexcer) { _ = "STUB: not implemented"; return }

func testDlaexc(t *testing.T, impl Dlaexcer, rnd *rand.Rand, n, extra int) {
	_ = "STUB: not implemented"
	return
}
