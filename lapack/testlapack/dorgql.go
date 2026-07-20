package testlapack

import (
	"testing"
)

type Dorgqler interface {
	Dorgql(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)

	Dlarfger
}

func DorgqlTest(t *testing.T, impl Dorgqler) { _ = "STUB: not implemented"; return }
