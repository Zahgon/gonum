package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dlapmrer interface {
	Dlapmr(forwrd bool, m, n int, x []float64, ldx int, k []int)
}

func DlapmrTest(t *testing.T, impl Dlapmrer) { _ = "STUB: not implemented"; return }

func dlapmrTest(t *testing.T, impl Dlapmrer, rnd *rand.Rand, fwd bool, m, n, ldx int) {
	_ = "STUB: not implemented"
	return
}
