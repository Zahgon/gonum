package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlarfber interface {
	Dlarfter
	Dlarfb(side blas.Side, trans blas.Transpose, direct lapack.Direct,
		store lapack.StoreV, m, n, k int, v []float64, ldv int, t []float64, ldt int,
		c []float64, ldc int, work []float64, ldwork int)
}

func DlarfbTest(t *testing.T, impl Dlarfber) { _ = "STUB: not implemented"; return }
