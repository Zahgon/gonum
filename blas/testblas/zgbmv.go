package testblas

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zgbmver interface {
	Zgbmv(trans blas.Transpose, m, n, kL, kU int, alpha complex128, ab []complex128, ldab int, x []complex128, incX int, beta complex128, y []complex128, incY int)

	Zgemver
}

func ZgbmvTest(t *testing.T, impl Zgbmver) { _ = "STUB: not implemented"; return }

func testZgbmv(t *testing.T, impl Zgbmver, rnd *rand.Rand, trans blas.Transpose, m, n, kL, kU int, alpha, beta complex128, ldab, incX, incY int) {
	_ = "STUB: not implemented"
	return
}
