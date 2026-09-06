package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dlatrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, normin bool, n int, a []float64, lda int, x []float64, cnorm []float64) (scale float64) {
	_ = "STUB: not implemented"
	return 0
}
