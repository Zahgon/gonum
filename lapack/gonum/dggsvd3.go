package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dggsvd3(jobU, jobV, jobQ lapack.GSVDJob, m, n, p int, a []float64, lda int, b []float64, ldb int, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64, lwork int, iwork []int) (k, l int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}
