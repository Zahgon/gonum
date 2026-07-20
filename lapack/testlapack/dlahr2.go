package testlapack

import (
	"testing"
)

type Dlahr2er interface {
	Dlahr2(n, k, nb int, a []float64, lda int, tau, t []float64, ldt int, y []float64, ldy int)
}

func Dlahr2Test(t *testing.T, impl Dlahr2er) { _ = "STUB: not implemented"; return }
