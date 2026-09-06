package testlapack

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
)

func Dlatm1(dst []float64, mode int, cond float64, rsign bool, dist int, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}

func Dlagsy(n, k int, d []float64, a []float64, lda int, rnd *rand.Rand, work []float64) {
	_ = "STUB: not implemented"
	return
}

func Dlagge(m, n, kl, ku int, d []float64, a []float64, lda int, rnd *rand.Rand, work []float64) {
	_ = "STUB: not implemented"
	return
}

func dlarnv(dst []float64, dist int, rnd *rand.Rand) { _ = "STUB: not implemented"; return }

func dlattr(imat int, uplo blas.Uplo, trans blas.Transpose, n int, a []float64, lda int, b, work []float64, rnd *rand.Rand) (diag blas.Diag) {
	_ = "STUB: not implemented"
	return *new(blas.Diag)
}

func checkMatrix(m, n int, a []float64, lda int) { _ = "STUB: not implemented"; return }

func randomOrthogonal(n int, rnd *rand.Rand) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func reflector(v, col []float64, j int) { _ = "STUB: not implemented"; return }

func applyReflector(qh blas64.General, q blas64.General, v []float64) {
	_ = "STUB: not implemented"
	return
}

func dlattb(kind int, uplo blas.Uplo, trans blas.Transpose, n, kd int, ab []float64, ldab int, rnd *rand.Rand) (diag blas.Diag, b []float64) {
	_ = "STUB: not implemented"
	return *new(blas.Diag), nil
}
