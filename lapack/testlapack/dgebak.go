package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

type Dgebaker interface {
	Dgebak(job lapack.BalanceJob, side lapack.EVSide, n, ilo, ihi int, scale []float64, m int, v []float64, ldv int)
}

func DgebakTest(t *testing.T, impl Dgebaker) { _ = "STUB: not implemented"; return }

func testDgebak(t *testing.T, impl Dgebaker, job lapack.BalanceJob, side lapack.EVSide, ilo, ihi int, v blas64.General, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
