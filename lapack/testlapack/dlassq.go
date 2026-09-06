package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dlassqer interface {
	Dlassq(n int, x []float64, incx int, scale, ssq float64) (float64, float64)
}

func DlassqTest(t *testing.T, impl Dlassqer) { _ = "STUB: not implemented"; return }

func dlassqTest(t *testing.T, impl Dlassqer, rnd *rand.Rand, n, incx, cas int, v0, v1 float64) {
	_ = "STUB: not implemented"
	return
}
