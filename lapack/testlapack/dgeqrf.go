package testlapack

import (
	"testing"
)

type Dgeqrfer interface {
	Dgeqr2er
	Dgeqrf(m, n int, a []float64, lda int, tau, work []float64, lwork int)
}

func DgeqrfTest(t *testing.T, impl Dgeqrfer) { _ = "STUB: not implemented"; return }
