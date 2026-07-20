package gonum

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

func (Implementation) Dlarfb(side blas.Side, trans blas.Transpose, direct lapack.Direct, store lapack.StoreV, m, n, k int, v []float64, ldv int, t []float64, ldt int, c []float64, ldc int, work []float64, ldwork int) {
	_ = "STUB: not implemented"
	return
}
