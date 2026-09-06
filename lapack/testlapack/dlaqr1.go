package testlapack

import (
	"testing"
)

type Dlaqr1er interface {
	Dlaqr1(n int, h []float64, ldh int, sr1, si1, sr2, si2 float64, v []float64)
}

func Dlaqr1Test(t *testing.T, impl Dlaqr1er) { _ = "STUB: not implemented"; return }
