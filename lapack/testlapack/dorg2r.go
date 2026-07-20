package testlapack

import (
	"testing"
)

type Dorg2rer interface {
	Dgeqrfer
	Dorg2r(m, n, k int, a []float64, lda int, tau []float64, work []float64)
}

func Dorg2rTest(t *testing.T, impl Dorg2rer) { _ = "STUB: not implemented"; return }
