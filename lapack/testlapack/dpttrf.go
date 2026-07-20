package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dpttrfer interface {
	Dpttrf(n int, d, e []float64) (ok bool)
}

func DpttrfTest(t *testing.T, impl Dpttrfer) { _ = "STUB: not implemented"; return }

func dpttrfTest(t *testing.T, impl Dpttrfer, rnd *rand.Rand, n int) {
	_ = "STUB: not implemented"
	return
}

func dpttrfResidual(n int, d, e, dFac, eFac []float64) float64 { _ = "STUB: not implemented"; return 0 }

func newRandomSymTridiag(n int, rnd *rand.Rand) (d, e []float64) {
	_ = "STUB: not implemented"
	return nil, nil
}
