package testlapack

import (
	"testing"
)

type Dgetrier interface {
	Dgetrfer
	Dgetri(n int, a []float64, lda int, ipiv []int, work []float64, lwork int) bool
}

func DgetriTest(t *testing.T, impl Dgetrier) { _ = "STUB: not implemented"; return }
