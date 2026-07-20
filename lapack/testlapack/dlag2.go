package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas/blas64"
)

type Dlag2er interface {
	Dlag2(a []float64, lda int, b []float64, ldb int) (scale1, scale2, wr1, wr2, wi float64)
}

func Dlag2Test(t *testing.T, impl Dlag2er) { _ = "STUB: not implemented"; return }

func dlag2Test(t *testing.T, impl Dlag2er, rnd *rand.Rand, lda, ldb int, aKind, bKind int) {
	_ = "STUB: not implemented"
	return
}

func makeDlag2TestMatrix(rnd *rand.Rand, ld, kind int) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func residualDlag2(a, b blas64.General, s float64, w complex128) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func zabs(z complex128) float64 { _ = "STUB: not implemented"; return 0 }

func scale(f float64, c complex128) complex128 { _ = "STUB: not implemented"; return 0 }

func cmplxdet2x2(a11, a12, a21, a22 complex128) complex128 { _ = "STUB: not implemented"; return 0 }
