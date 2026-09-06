package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpotrier interface {
	Dpotri(uplo blas.Uplo, n int, a []float64, lda int) bool

	Dpotrf(uplo blas.Uplo, n int, a []float64, lda int) bool
}

func DpotriTest(t *testing.T, impl Dpotrier) { _ = "STUB: not implemented"; return }
