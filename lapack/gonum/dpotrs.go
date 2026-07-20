package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (Implementation) Dpotrs(uplo blas.Uplo, n, nrhs int, a []float64, lda int, b []float64, ldb int) {
	_ = "STUB: not implemented"
	return
}
