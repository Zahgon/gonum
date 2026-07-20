package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dlascl(kind lapack.MatrixType, kl, ku int, cfrom, cto float64, m, n int, a []float64, lda int) {
	_ = "STUB: not implemented"
	return
}
