package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpbtrser interface {
	Dpbtrs(uplo blas.Uplo, n, kd, nrhs int, ab []float64, ldab int, b []float64, ldb int)

	Dpbtrfer
}

func DpbtrsTest(t *testing.T, impl Dpbtrser) { _ = "STUB: not implemented"; return }

func dpbtrsTest(t *testing.T, impl Dpbtrser, rnd *rand.Rand, uplo blas.Uplo, n, kd, nrhs int, ldab, ldb int) {
	_ = "STUB: not implemented"
	return
}
