package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (Implementation) Dpbtrs(uplo blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int) {
	_ = "STUB: not implemented"
	return
}
