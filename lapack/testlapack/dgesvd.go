package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dgesvder interface {
	Dgesvd(jobU, jobVT lapack.SVDJob, m, n int, a []float64, lda int, s, u []float64, ldu int, vt []float64, ldvt int, work []float64, lwork int) (ok bool)
}

func DgesvdTest(t *testing.T, impl Dgesvder, tol float64) { _ = "STUB: not implemented"; return }

func dgesvdTest(t *testing.T, impl Dgesvder, m, n, mtype int, tol float64) {
	_ = "STUB: not implemented"
	return
}

func svdFullResidual(m, n int, aNorm float64, a []float64, lda int, u []float64, ldu int, d []float64, vt []float64, ldvt int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func svdPartialUResidual(m, n int, u, uRef []float64, ldu int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func svdPartialVTResidual(m, n int, vt, vtRef []float64, ldvt int) float64 {
	_ = "STUB: not implemented"
	return 0
}
