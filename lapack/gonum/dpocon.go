package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dpocon(uplo blas.Uplo, n int, a []float64, lda int, anorm float64, work []float64, iwork []int) float64 {
	_ = "STUB: not implemented"
	return 0
}
