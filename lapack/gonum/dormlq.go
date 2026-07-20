package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dormlq(side blas.Side, trans blas.Transpose, m, n, k int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int) {
	_ = "STUB: not implemented"
	return
}
