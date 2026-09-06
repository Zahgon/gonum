package testlapack

import (
	"testing"
)

type Dorgqrer interface {
	Dorg2rer
	Dorgqr(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
}

func DorgqrTest(t *testing.T, impl Dorgqrer) { _ = "STUB: not implemented"; return }
