package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dgetrs(trans blas.Transpose, n, nrhs int, a []float64, lda int, ipiv []int, b []float64, ldb int) {
	_ = "STUB: not implemented"
	return
}
