package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpotrser interface {
	Dpotrs(uplo blas.Uplo, n, nrhs int, a []float64, lda int, b []float64, ldb int)

	Dpotrf(uplo blas.Uplo, n int, a []float64, lda int) bool
}

func DpotrsTest(t *testing.T, impl Dpotrser) { _ = "STUB: not implemented"; return }
