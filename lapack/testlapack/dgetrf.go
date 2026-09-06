package testlapack

import (
	"testing"
)

type Dgetrfer interface {
	Dgetrf(m, n int, a []float64, lda int, ipiv []int) bool
}

func DgetrfTest(t *testing.T, impl Dgetrfer) { _ = "STUB: not implemented"; return }
