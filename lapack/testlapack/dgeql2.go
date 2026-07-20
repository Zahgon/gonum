package testlapack

import (
	"testing"
)

type Dgeql2er interface {
	Dgeql2(m, n int, a []float64, lda int, tau, work []float64)
}

func Dgeql2Test(t *testing.T, impl Dgeql2er) { _ = "STUB: not implemented"; return }
