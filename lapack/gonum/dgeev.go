package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dgeev(jobvl lapack.LeftEVJob, jobvr lapack.RightEVJob, n int, a []float64, lda int, wr, wi []float64, vl []float64, ldvl int, vr []float64, ldvr int, work []float64, lwork int) (first int) {
	_ = "STUB: not implemented"
	return 0
}
