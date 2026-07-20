package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrti2er interface {
	Dtrti2(uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int)
}

func Dtrti2Test(t *testing.T, impl Dtrti2er) { _ = "STUB: not implemented"; return }
