package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dgghrd(compq, compz lapack.OrthoComp, n, ilo, ihi int, a []float64, lda int, b []float64, ldb int, q []float64, ldq int, z []float64, ldz int) {
	_ = "STUB: not implemented"
	return
}
