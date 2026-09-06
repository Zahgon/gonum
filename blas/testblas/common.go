package testblas

import (
	"math/cmplx"
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

const throwPanic = true

var znan = cmplx.NaN()

func dTolEqual(a, b float64) bool { _ = "STUB: not implemented"; return false }

func dSliceTolEqual(a, b []float64) bool { _ = "STUB: not implemented"; return false }

func dStridedSliceTolEqual(n int, a []float64, inca int, b []float64, incb int) bool {
	_ = "STUB: not implemented"
	return false
}

func dSliceEqual(a, b []float64) bool { _ = "STUB: not implemented"; return false }

func dCopyTwoTmp(x, xTmp, y, yTmp []float64) { _ = "STUB: not implemented"; return }

func panics(f func()) (b bool) { _ = "STUB: not implemented"; return false }

func testpanics(f func(), name string, t *testing.T) { _ = "STUB: not implemented"; return }

func sliceOfSliceCopy(a [][]float64) [][]float64 { _ = "STUB: not implemented"; return nil }

func sliceCopy(a []float64) []float64 { _ = "STUB: not implemented"; return nil }

func flatten(a [][]float64) []float64 { _ = "STUB: not implemented"; return nil }

func unflatten(a []float64, m, n int) [][]float64 { _ = "STUB: not implemented"; return nil }

func flattenTriangular(a [][]float64, ul blas.Uplo) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func flattenBanded(a [][]float64, ku, kl int) []float64 { _ = "STUB: not implemented"; return nil }

func makeIncremented(x []float64, inc int, extra int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func makeIncremented32(x []float32, inc int, extra int) []float32 {
	_ = "STUB: not implemented"
	return nil
}

func abs(x int) int { _ = "STUB: not implemented"; return 0 }

func allPairs(x, y []int) [][2]int { _ = "STUB: not implemented"; return nil }

func sameFloat64(a, b float64) bool { _ = "STUB: not implemented"; return false }

func sameComplex128(x, y complex128) bool { _ = "STUB: not implemented"; return false }

func zsame(x, y []complex128) bool { _ = "STUB: not implemented"; return false }

func zSameAtNonstrided(x, y []complex128, inc int) bool { _ = "STUB: not implemented"; return false }

func zEqualApproxAtStrided(x, y []complex128, inc int, tol float64) bool {
	_ = "STUB: not implemented"
	return false
}

func makeZVector(data []complex128, inc int) []complex128 { _ = "STUB: not implemented"; return nil }

func makeZGeneral(data []complex128, m, n int, ld int) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func zPack(uplo blas.Uplo, n int, a []complex128, lda int) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func zUnpackAsHermitian(uplo blas.Uplo, n int, ap []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func zPackBand(kL, kU, ldab int, m, n int, a []complex128, lda int) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func zPackTriBand(k, ldab int, uplo blas.Uplo, n int, a []complex128, lda int) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func zEqualApprox(a, b []complex128, tol float64) bool { _ = "STUB: not implemented"; return false }

func rndComplex128(rnd *rand.Rand) complex128 { _ = "STUB: not implemented"; return 0 }

func zmm(tA, tB blas.Transpose, m, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func transString(t blas.Transpose) string { _ = "STUB: not implemented"; return "" }

func uploString(uplo blas.Uplo) string { _ = "STUB: not implemented"; return "" }

func sideString(side blas.Side) string { _ = "STUB: not implemented"; return "" }

func diagString(diag blas.Diag) string { _ = "STUB: not implemented"; return "" }

func zSameLowerTri(n int, a []complex128, lda int, b []complex128, ldb int) bool {
	_ = "STUB: not implemented"
	return false
}

func zSameUpperTri(n int, a []complex128, lda int, b []complex128, ldb int) bool {
	_ = "STUB: not implemented"
	return false
}
