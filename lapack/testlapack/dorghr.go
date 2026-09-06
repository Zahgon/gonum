package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dorghrer interface {
	Dorghr(n, ilo, ihi int, a []float64, lda int, tau, work []float64, lwork int)

	Dgehrder
}

func DorghrTest(t *testing.T, impl Dorghrer) { _ = "STUB: not implemented"; return }

func testDorghr(t *testing.T, impl Dorghrer, n, ilo, ihi, extra int, optwork bool, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
