package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dggsvp3(jobU, jobV, jobQ lapack.GSVDJob, m, p, n int, a []float64, lda int, b []float64, ldb int, tola, tolb float64, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, iwork []int, tau, work []float64, lwork int) (k, l int) {
	_ = "STUB: not implemented"
	return 0, 0
}
