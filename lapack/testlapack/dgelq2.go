package testlapack

import (
	"testing"
)

type Dgelq2er interface {
	Dgelq2(m, n int, a []float64, lda int, tau, work []float64)
}

func Dgelq2Test(t *testing.T, impl Dgelq2er) { _ = "STUB: not implemented"; return }
