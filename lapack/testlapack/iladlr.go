package testlapack

import "testing"

type Iladlrer interface {
	Iladlr(m, n int, a []float64, lda int) int
}

func IladlrTest(t *testing.T, impl Iladlrer) { _ = "STUB: not implemented"; return }
