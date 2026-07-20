package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dptsver interface {
	Dptsv(n, nrhs int, d, e []float64, b []float64, ldb int) (ok bool)
}

func DptsvTest(t *testing.T, impl Dptsver) { _ = "STUB: not implemented"; return }

func dptsvTest(t *testing.T, impl Dptsver, rnd *rand.Rand, n, nrhs, ldb int) {
	_ = "STUB: not implemented"
	return
}
