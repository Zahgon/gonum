package testlapack

import (
	"testing"
)

type Dlabrder interface {
	Dlabrd(m, n, nb int, a []float64, lda int, d, e, tauq, taup, x []float64, ldx int, y []float64, ldy int)
}

func DlabrdTest(t *testing.T, impl Dlabrder) { _ = "STUB: not implemented"; return }
