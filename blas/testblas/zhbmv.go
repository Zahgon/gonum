package testblas

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Zhbmver interface {
	Zhbmv(uplo blas.Uplo, n, k int, alpha complex128, ab []complex128, ldab int, x []complex128, incX int, beta complex128, y []complex128, incY int)

	Zhemver
}

func ZhbmvTest(t *testing.T, impl Zhbmver) { _ = "STUB: not implemented"; return }

func testZhbmv(t *testing.T, impl Zhbmver, rnd *rand.Rand, uplo blas.Uplo, n, k int, alpha, beta complex128, ldab, incX, incY int) {
	_ = "STUB: not implemented"
	return
}
