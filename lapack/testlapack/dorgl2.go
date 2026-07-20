package testlapack

import (
	"testing"
)

type Dorgl2er interface {
	Dgelqfer
	Dorgl2(m, n, k int, a []float64, lda int, tau []float64, work []float64)
}

func Dorgl2Test(t *testing.T, impl Dorgl2er) { _ = "STUB: not implemented"; return }
