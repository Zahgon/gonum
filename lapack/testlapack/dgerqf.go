package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgerqfer interface {
	Dgerqf(m, n int, a []float64, lda int, tau, work []float64, lwork int)
}

func DgerqfTest(t *testing.T, impl Dgerqfer) { _ = "STUB: not implemented"; return }

func dgerqfTest(t *testing.T, impl Dgerqfer, rnd *rand.Rand, m, n, lda int) {
	_ = "STUB: not implemented"
	return
}
