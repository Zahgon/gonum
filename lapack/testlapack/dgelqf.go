package testlapack

import (
	"testing"
)

type Dgelqfer interface {
	Dgelq2er
	Dgelqf(m, n int, a []float64, lda int, tau, work []float64, lwork int)
}

func DgelqfTest(t *testing.T, impl Dgelqfer) { _ = "STUB: not implemented"; return }
