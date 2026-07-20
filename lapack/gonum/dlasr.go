package gonum

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dlasr(side blas.Side, pivot lapack.Pivot, direct lapack.Direct, m, n int, c, s, a []float64, lda int) {
	_ = "STUB: not implemented"
	return
}
