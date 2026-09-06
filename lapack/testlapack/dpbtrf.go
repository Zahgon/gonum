package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpbtrfer interface {
	Dpbtrf(uplo blas.Uplo, n, kd int, ab []float64, ldab int) (ok bool)
}

func DpbtrfTest(t *testing.T, impl Dpbtrfer) { _ = "STUB: not implemented"; return }

func dpbtrfTest(t *testing.T, impl Dpbtrfer, uplo blas.Uplo, n, kd int, ldab int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}

func dsbmm(uplo blas.Uplo, n, kd int, ab []float64, ldab int) { _ = "STUB: not implemented"; return }
