package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (Implementation) Dlatbs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, normin bool, n, kd int, ab []float64, ldab int, x, cnorm []float64) (scale float64) {
	_ = "STUB: not implemented"
	return 0
}
