package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dbdsqr(uplo blas.Uplo, n, ncvt, nru, ncc int, d, e, vt []float64, ldvt int, u []float64, ldu int, c []float64, ldc int, work []float64) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
