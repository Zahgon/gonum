package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

func (impl Implementation) Dtgsja(jobU, jobV, jobQ lapack.GSVDJob, m, p, n, k, l int, a []float64, lda int, b []float64, ldb int, tola, tolb float64, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64) (cycles int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}
