package testlapack

import (
	"math/rand/v2"
	"testing"
)

type Dgesver interface {
	Dgesv(n, nrhs int, a []float64, lda int, ipiv []int, b []float64, ldb int) bool

	Dgetri(n int, a []float64, lda int, ipiv []int, work []float64, lwork int) bool
}

func DgesvTest(t *testing.T, impl Dgesver) { _ = "STUB: not implemented"; return }

func dgesvTest(t *testing.T, impl Dgesver, rnd *rand.Rand, n, nrhs, lda, ldb int) {
	_ = "STUB: not implemented"
	return
}
