package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (Implementation) Dlarft(direct lapack.Direct, store lapack.StoreV, n, k int, v []float64, ldv int, tau []float64, t []float64, ldt int) {
	_ = "STUB: not implemented"
	return
}
