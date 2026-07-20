package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dggsvd3er interface {
	Dggsvd3(jobU, jobV, jobQ lapack.GSVDJob, m, n, p int, a []float64, lda int, b []float64, ldb int, alpha, beta, u []float64, ldu int, v []float64, ldv int, q []float64, ldq int, work []float64, lwork int, iwork []int) (k, l int, ok bool)
}

func Dggsvd3Test(t *testing.T, impl Dggsvd3er) { _ = "STUB: not implemented"; return }
