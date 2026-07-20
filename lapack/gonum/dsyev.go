package gonum

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dsyev(jobz lapack.EVJob, uplo blas.Uplo, n int, a []float64, lda int, w, work []float64, lwork int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
