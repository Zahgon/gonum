package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dlangb(norm lapack.MatrixNorm, m, n, kl, ku int, ab []float64, ldab int) float64 {
	_ = "STUB: not implemented"
	return 0
}
