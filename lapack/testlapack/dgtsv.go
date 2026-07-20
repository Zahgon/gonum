package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgtsver interface {
	Dgtsv(n, nrhs int, dl, d, du []float64, b []float64, ldb int) (ok bool)
}

func DgtsvTest(t *testing.T, impl Dgtsver) { _ = "STUB: not implemented"; return }

func dgtsvTest(t *testing.T, impl Dgtsver, rnd *rand.Rand, n, nrhs, ldb int) {
	_ = "STUB: not implemented"
	return
}
