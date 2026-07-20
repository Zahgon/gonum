package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpotf2er interface {
	Dpotf2(ul blas.Uplo, n int, a []float64, lda int) (ok bool)
}

func Dpotf2Test(t *testing.T, impl Dpotf2er) { _ = "STUB: not implemented"; return }

func testDpotf2(t *testing.T, impl Dpotf2er, testPos bool, a, ans [][]float64, stride int, ul blas.Uplo) {
	_ = "STUB: not implemented"
	return
}

func flattenTri(a [][]float64, stride int, ul blas.Uplo) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func transpose(a [][]float64) [][]float64 { _ = "STUB: not implemented"; return nil }
