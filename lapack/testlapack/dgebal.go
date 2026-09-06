package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

type Dgebaler interface {
	Dgebal(job lapack.BalanceJob, n int, a []float64, lda int, scale []float64) (int, int)
}

func DgebalTest(t *testing.T, impl Dgebaler) { _ = "STUB: not implemented"; return }

func testDgebal(t *testing.T, impl Dgebaler, job lapack.BalanceJob, a blas64.General) {
	_ = "STUB: not implemented"
	return
}
