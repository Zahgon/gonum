package testlapack

import (
	"testing"
)

type Dlasq1er interface {
	Dlasq1(n int, d, e, work []float64) int

	Dgebrd(m, n int, a []float64, lda int, d, e, tauQ, tauP, work []float64, lwork int)
}

func Dlasq1Test(t *testing.T, impl Dlasq1er) { _ = "STUB: not implemented"; return }
