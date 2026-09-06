package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlauu2er interface {
	Dlauu2(uplo blas.Uplo, n int, a []float64, lda int)
}

func Dlauu2Test(t *testing.T, impl Dlauu2er) { _ = "STUB: not implemented"; return }

func dlauuTest(t *testing.T, dlauu func(blas.Uplo, int, []float64, int), uplo blas.Uplo, ns []int) {
	_ = "STUB: not implemented"
	return
}
