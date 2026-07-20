package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dgels(trans blas.Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int, work []float64, lwork int) bool {
	_ = "STUB: not implemented"
	return false
}
