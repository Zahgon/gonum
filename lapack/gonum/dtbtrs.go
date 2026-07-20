package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dtbtrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, kd, nrhs int, a []float64, lda int, b []float64, ldb int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
