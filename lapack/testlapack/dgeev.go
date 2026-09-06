package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

type Dgeever interface {
	Dgeev(jobvl lapack.LeftEVJob, jobvr lapack.RightEVJob, n int, a []float64, lda int,
		wr, wi []float64, vl []float64, ldvl int, vr []float64, ldvr int, work []float64, lwork int) int
}

type dgeevTest struct {
	a      blas64.General
	evWant []complex128
	valTol float64
	vecTol float64
}

func DgeevTest(t *testing.T, impl Dgeever) { _ = "STUB: not implemented"; return }

func testDgeev(t *testing.T, impl Dgeever, tc string, test dgeevTest, jobvl lapack.LeftEVJob, jobvr lapack.RightEVJob, extra int, wl worklen) {
	_ = "STUB: not implemented"
	return
}

func dgeevTestForAntisymRandom(n int, rnd *rand.Rand) dgeevTest {
	_ = "STUB: not implemented"
	return *new(dgeevTest)
}

func residualRightEV(a, e blas64.General, wr, wi []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func residualLeftEV(a, e blas64.General, wr, wi []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
