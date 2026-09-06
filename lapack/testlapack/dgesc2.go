package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgesc2er interface {
	Dgesc2(n int, a []float64, lda int, rhs []float64, ipiv, jpiv []int) (scale float64)

	Dgetc2er
}

func Dgesc2Test(t *testing.T, impl Dgesc2er) { _ = "STUB: not implemented"; return }

func testDgesc2(t *testing.T, impl Dgesc2er, rnd *rand.Rand, n, lda int, big bool) {
	_ = "STUB: not implemented"
	return
}
