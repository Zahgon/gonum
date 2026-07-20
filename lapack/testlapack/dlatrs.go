package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlatrser interface {
	Dlatrs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, normin bool, n int, a []float64, lda int, x []float64, cnorm []float64) (scale float64)
}

func DlatrsTest(t *testing.T, impl Dlatrser) { _ = "STUB: not implemented"; return }

func testDlatrs(t *testing.T, impl Dlatrser, imat int, uplo blas.Uplo, trans blas.Transpose, n, lda int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}

func dlatrsResidual(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, a []float64, lda int, scale float64, cnorm []float64, x, b, work []float64) (resid float64, hasNaN bool) {
	_ = "STUB: not implemented"
	return 0, false
}
