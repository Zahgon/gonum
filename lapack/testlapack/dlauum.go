package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlauumer interface {
	Dlauum(uplo blas.Uplo, n int, a []float64, lda int)
}

func DlauumTest(t *testing.T, impl Dlauumer) { _ = "STUB: not implemented"; return }
