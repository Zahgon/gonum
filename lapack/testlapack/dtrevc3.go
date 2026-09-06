package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

type Dtrevc3er interface {
	Dtrevc3(side lapack.EVSide, howmny lapack.EVHowMany, selected []bool, n int, t []float64, ldt int, vl []float64, ldvl int, vr []float64, ldvr int, mm int, work []float64, lwork int) int
}

func Dtrevc3Test(t *testing.T, impl Dtrevc3er) { _ = "STUB: not implemented"; return }

func runDtrevc3Test(t *testing.T, impl Dtrevc3er, rnd *rand.Rand, side lapack.EVSide) {
	_ = "STUB: not implemented"
	return
}

func dtrevc3Test(t *testing.T, impl Dtrevc3er, side lapack.EVSide, n, extra int, optwork bool, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}

func residualEVNormalization(emat blas64.General, wi []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func normalizeEV(emat blas64.General, wi []float64) { _ = "STUB: not implemented"; return }
