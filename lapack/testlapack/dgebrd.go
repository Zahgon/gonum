package testlapack

import (
	"testing"
)

type Dgebrder interface {
	Dgebrd(m, n int, a []float64, lda int, d, e, tauQ, tauP, work []float64, lwork int)
	Dgebd2er
}

func DgebrdTest(t *testing.T, impl Dgebrder) { _ = "STUB: not implemented"; return }
