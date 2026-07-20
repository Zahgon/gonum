package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (impl Implementation) Dpbcon(uplo blas.Uplo, n, kd int, ab []float64, ldab int, anorm float64, work []float64, iwork []int) (rcond float64) {
	_ = "STUB: not implemented"
	return 0
}
