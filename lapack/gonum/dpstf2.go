package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (Implementation) Dpstf2(uplo blas.Uplo, n int, a []float64, lda int, piv []int, tol float64, work []float64) (rank int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}
