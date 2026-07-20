package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgehrder interface {
	Dgehrd(n, ilo, ihi int, a []float64, lda int, tau, work []float64, lwork int)

	Dorgqr(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
}

func DgehrdTest(t *testing.T, impl Dgehrder) { _ = "STUB: not implemented"; return }

func testDgehrd(t *testing.T, impl Dgehrder, n, ilo, ihi, extra int, optwork bool, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
