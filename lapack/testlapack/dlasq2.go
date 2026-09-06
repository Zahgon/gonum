package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlasq2er interface {
	Dlasq2(n int, z []float64) (info int)

	Dsyev(jobz lapack.EVJob, uplo blas.Uplo, n int, a []float64, lda int, w, work []float64, lwork int) (ok bool)
}

func Dlasq2Test(t *testing.T, impl Dlasq2er) { _ = "STUB: not implemented"; return }
