package testlapack

import "testing"

type Iladlcer interface {
	Iladlc(m, n int, a []float64, lda int) int
}

func IladlcTest(t *testing.T, impl Iladlcer) { _ = "STUB: not implemented"; return }
