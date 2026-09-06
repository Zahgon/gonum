package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsytd2er interface {
	Dsytd2(uplo blas.Uplo, n int, a []float64, lda int, d, e, tau []float64)
}

func Dsytd2Test(t *testing.T, impl Dsytd2er) { _ = "STUB: not implemented"; return }
