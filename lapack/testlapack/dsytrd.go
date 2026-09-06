package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dsytrder interface {
	Dsytrd(uplo blas.Uplo, n int, a []float64, lda int, d, e, tau, work []float64, lwork int)

	Dorgqr(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
	Dorgql(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
}

func DsytrdTest(t *testing.T, impl Dsytrder) { _ = "STUB: not implemented"; return }
