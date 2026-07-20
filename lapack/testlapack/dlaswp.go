package testlapack

import (
	"testing"
)

type Dlaswper interface {
	Dlaswp(n int, a []float64, lda, k1, k2 int, ipiv []int, incX int)
}

func DlaswpTest(t *testing.T, impl Dlaswper) { _ = "STUB: not implemented"; return }
