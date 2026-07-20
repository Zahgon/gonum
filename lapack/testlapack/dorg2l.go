package testlapack

import (
	"testing"
)

type Dorg2ler interface {
	Dorg2l(m, n, k int, a []float64, lda int, tau, work []float64)
	Dgeql2er
}

func Dorg2lTest(t *testing.T, impl Dorg2ler) { _ = "STUB: not implemented"; return }
