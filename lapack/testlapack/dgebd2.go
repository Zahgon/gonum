package testlapack

import (
	"testing"
)

type Dgebd2er interface {
	Dgebd2(m, n int, a []float64, lda int, d, e, tauq, taup, work []float64)
}

func Dgebd2Test(t *testing.T, impl Dgebd2er) { _ = "STUB: not implemented"; return }
