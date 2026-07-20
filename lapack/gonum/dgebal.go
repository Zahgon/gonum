package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dgebal(job lapack.BalanceJob, n int, a []float64, lda int, scale []float64) (ilo, ihi int) {
	_ = "STUB: not implemented"
	return 0, 0
}
