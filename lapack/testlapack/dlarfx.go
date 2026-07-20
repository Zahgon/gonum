package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlarfxer interface {
	Dlarfx(side blas.Side, m, n int, v []float64, tau float64, c []float64, ldc int, work []float64)
}

func DlarfxTest(t *testing.T, impl Dlarfxer) { _ = "STUB: not implemented"; return }

func testDlarfx(t *testing.T, impl Dlarfxer, side blas.Side, m, n, extra int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
