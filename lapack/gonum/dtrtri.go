package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dtrtri(uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
