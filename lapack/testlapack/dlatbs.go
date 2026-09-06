package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dlatbser interface {
	Dlatbs(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, normin bool, n, kd int, ab []float64, ldab int, x []float64, cnorm []float64) float64
}

func DlatbsTest(t *testing.T, impl Dlatbser) { _ = "STUB: not implemented"; return }

func dlatbsTest(t *testing.T, impl Dlatbser, rnd *rand.Rand, kind int, uplo blas.Uplo, trans blas.Transpose, n, kd, ldab int) {
	_ = "STUB: not implemented"
	return
}

func dlatbsResidual(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, kd int, ab []float64, ldab int, scale float64, cnorm, b, x []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
