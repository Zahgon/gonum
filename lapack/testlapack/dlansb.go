package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlansber interface {
	Dlansb(norm lapack.MatrixNorm, uplo blas.Uplo, n, kd int, ab []float64, ldab int, work []float64) float64
}

func DlansbTest(t *testing.T, impl Dlansber) { _ = "STUB: not implemented"; return }

func dlansbTest(t *testing.T, impl Dlansber, rnd *rand.Rand, uplo blas.Uplo, n, kd int, ldab int) {
	_ = "STUB: not implemented"
	return
}
