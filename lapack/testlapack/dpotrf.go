package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dpotrfer interface {
	Dpotrf(ul blas.Uplo, n int, a []float64, lda int) (ok bool)
}

func DpotrfTest(t *testing.T, impl Dpotrfer) { _ = "STUB: not implemented"; return }
