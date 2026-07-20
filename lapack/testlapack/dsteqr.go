package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

type Dsteqrer interface {
	Dsteqr(compz lapack.EVComp, n int, d, e, z []float64, ldz int, work []float64) (ok bool)
	Dorgtrer
}

func DsteqrTest(t *testing.T, impl Dsteqrer) { _ = "STUB: not implemented"; return }

func eigenDecompCorrect(values []float64, A, V blas64.General) bool {
	_ = "STUB: not implemented"
	return false
}
