package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlacpyer interface {
	Dlacpy(uplo blas.Uplo, m, n int, a []float64, lda int, b []float64, ldb int)
}

func DlacpyTest(t *testing.T, impl Dlacpyer) { _ = "STUB: not implemented"; return }
