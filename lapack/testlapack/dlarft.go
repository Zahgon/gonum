package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlarfter interface {
	Dgeqr2er
	Dlarft(direct lapack.Direct, store lapack.StoreV, n, k int, v []float64, ldv int, tau []float64, t []float64, ldt int)
}

func DlarftTest(t *testing.T, impl Dlarfter) { _ = "STUB: not implemented"; return }
