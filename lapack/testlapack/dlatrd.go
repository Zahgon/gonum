package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

type Dlatrder interface {
	Dlatrd(uplo blas.Uplo, n, nb int, a []float64, lda int, e, tau, w []float64, ldw int)
}

func DlatrdTest(t *testing.T, impl Dlatrder) { _ = "STUB: not implemented"; return }

func dlatrdCheckDecomposition(t *testing.T, uplo blas.Uplo, n, nb int, e, a []float64, lda int, aGen, q blas64.General, tol float64) bool {
	_ = "STUB: not implemented"
	return false
}

func genFromSym(a blas64.Symmetric) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}
