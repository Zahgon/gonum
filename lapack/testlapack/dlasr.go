package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlasrer interface {
	Dlasr(side blas.Side, pivot lapack.Pivot, direct lapack.Direct, m, n int, c, s, a []float64, lda int)
}

func DlasrTest(t *testing.T, impl Dlasrer) { _ = "STUB: not implemented"; return }
