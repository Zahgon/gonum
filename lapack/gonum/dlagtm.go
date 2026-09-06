package gonum

import "gonum.org/v1/gonum/blas"

func (impl Implementation) Dlagtm(trans blas.Transpose, m, n int, alpha float64, dl, d, du []float64, b []float64, ldb int, beta float64, c []float64, ldc int) {
	_ = "STUB: not implemented"
	return
}
