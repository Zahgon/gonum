package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dsyever interface {
	Dsyev(jobz lapack.EVJob, uplo blas.Uplo, n int, a []float64, lda int, w, work []float64, lwork int) (ok bool)
}

func DsyevTest(t *testing.T, impl Dsyever) { _ = "STUB: not implemented"; return }
