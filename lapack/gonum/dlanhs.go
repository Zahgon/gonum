package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dlanhs(norm lapack.MatrixNorm, n int, a []float64, lda int, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
