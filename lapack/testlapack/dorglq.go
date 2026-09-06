package testlapack

import (
	"testing"
)

type Dorglqer interface {
	Dorgl2er
	Dorglq(m, n, k int, a []float64, lda int, tau, work []float64, lwork int)
}

func DorglqTest(t *testing.T, impl Dorglqer) { _ = "STUB: not implemented"; return }
