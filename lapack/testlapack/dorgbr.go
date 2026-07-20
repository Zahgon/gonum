package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dorgbrer interface {
	Dorgbr(vect lapack.GenOrtho, m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
	Dgebrder
}

func DorgbrTest(t *testing.T, impl Dorgbrer) { _ = "STUB: not implemented"; return }
