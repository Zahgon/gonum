package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlaseter interface {
	Dlaset(uplo blas.Uplo, m, n int, alpha, beta float64, a []float64, lda int)
}

func DlasetTest(t *testing.T, impl Dlaseter) { _ = "STUB: not implemented"; return }
