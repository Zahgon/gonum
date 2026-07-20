package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dormhrer interface {
	Dormhr(side blas.Side, trans blas.Transpose, m, n, ilo, ihi int, a []float64, lda int, tau, c []float64, ldc int, work []float64, lwork int)

	Dgehrder
}

func DormhrTest(t *testing.T, impl Dormhrer) { _ = "STUB: not implemented"; return }

func testDormhr(t *testing.T, impl Dormhrer, side blas.Side, trans blas.Transpose, m, n, ilo, ihi, extra int, optwork bool, rnd *rand.Rand) {
	_ = "STUB: not implemented"
	return
}
