package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpbconer interface {
	Dpbcon(uplo blas.Uplo, n, kd int, ab []float64, ldab int, anorm float64, work []float64, iwork []int) float64

	Dpbtrser
}

func DpbconTest(t *testing.T, impl Dpbconer) { _ = "STUB: not implemented"; return }

func dpbconTest(t *testing.T, impl Dpbconer, uplo blas.Uplo, n, kd, ldab int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}

func rCondTestRatio(rcond, rcondc float64) float64 { _ = "STUB: not implemented"; return 0 }
