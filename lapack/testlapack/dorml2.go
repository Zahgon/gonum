package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dorml2er interface {
	Dgelqfer
	Dorml2(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64)
}

func Dorml2Test(t *testing.T, impl Dorml2er) { _ = "STUB: not implemented"; return }
