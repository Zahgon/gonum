package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dsteqr(compz lapack.EVComp, n int, d, e, z []float64, ldz int, work []float64) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
