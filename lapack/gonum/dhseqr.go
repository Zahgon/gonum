package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dhseqr(job lapack.SchurJob, compz lapack.SchurComp, n, ilo, ihi int, h []float64, ldh int, wr, wi []float64, z []float64, ldz int, work []float64, lwork int) (unconverged int) {
	_ = "STUB: not implemented"
	return 0
}
