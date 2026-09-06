package gonum

import (
	"gonum.org/v1/gonum/lapack"
)

const noSVDO = "dgesvd: not coded for overwrite"

func (impl Implementation) Dgesvd(jobU, jobVT lapack.SVDJob, m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}
