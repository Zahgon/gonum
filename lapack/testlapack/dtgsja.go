package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dtgsjaer interface {
	Dlanger
	Dtgsja(jobU, jobV, jobQ lapack.GSVDJob, m, p, n, k, l int, a []float64, lda int, b []float64, ldb int, tola, tolb float64, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64) (cycles int, ok bool)
}

func DtgsjaTest(t *testing.T, impl Dtgsjaer) { _ = "STUB: not implemented"; return }
