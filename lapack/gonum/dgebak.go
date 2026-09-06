package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dgebak(job lapack.BalanceJob, side lapack.EVSide, n, ilo, ihi int, scale []float64, m int, v []float64, ldv int) {
	_ = "STUB: not implemented"
	return
}
