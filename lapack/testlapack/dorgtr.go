package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dorgtrer interface {
	Dorgtr(uplo blas.Uplo, n int, a []float64, lda int, tau, work []float64, lwork int)
	Dsytrder
}

func DorgtrTest(t *testing.T, impl Dorgtrer) { _ = "STUB: not implemented"; return }
