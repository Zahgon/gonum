package gonum

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dlansb(norm lapack.MatrixNorm, uplo blas.Uplo, n, kd int, ab []float64, ldab int, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
