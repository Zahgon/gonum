package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpbtf2er interface {
	Dpbtf2(uplo blas.Uplo, n, kd int, ab []float64, ldab int) (ok bool)
}

func Dpbtf2Test(t *testing.T, impl Dpbtf2er) { _ = "STUB: not implemented"; return }

func dpbtf2Test(t *testing.T, impl Dpbtf2er, rnd *rand.Rand, uplo blas.Uplo, n, kd int, ldab int) {
	_ = "STUB: not implemented"
	return
}
