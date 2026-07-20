package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dggsvp3er interface {
	Dlanger
	Dggsvp3(jobU, jobV, jobQ lapack.GSVDJob, m, p, n int, a []float64, lda int, b []float64, ldb int, tola, tolb float64, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, iwork []int, tau, work []float64, lwork int) (k, l int)
}

func Dggsvp3Test(t *testing.T, impl Dggsvp3er) { _ = "STUB: not implemented"; return }
