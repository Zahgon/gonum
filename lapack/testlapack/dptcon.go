package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dptconer interface {
	Dptcon(n int, d, e []float64, anorm float64, work []float64) (rcond float64)

	Dpttrf(n int, d, e []float64) (ok bool)
	Dpttrs(n, nrhs int, d, e []float64, b []float64, ldb int)
}

func DptconTest(t *testing.T, impl Dptconer) { _ = "STUB: not implemented"; return }

func dptconTest(t *testing.T, impl Dptconer, rnd *rand.Rand, n int) {
	_ = "STUB: not implemented"
	return
}
