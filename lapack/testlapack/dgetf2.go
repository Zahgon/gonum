package testlapack

import (
	"testing"
)

type Dgetf2er interface {
	Dgetf2(m, n int, a []float64, lda int, ipiv []int) bool
}

func Dgetf2Test(t *testing.T, impl Dgetf2er) { _ = "STUB: not implemented"; return }

func checkPLU(t *testing.T, ok bool, m, n, lda int, ipiv []int, factorized, original []float64, tol float64, print bool) {
	_ = "STUB: not implemented"
	return
}
