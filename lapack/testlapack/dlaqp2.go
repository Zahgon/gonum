package testlapack

import (
	"testing"
)

type Dlaqp2er interface {
	Dlapmter
	Dlaqp2(m, n, offset int, a []float64, lda int, jpvt []int, tau, vn1, vn2, work []float64)
}

func Dlaqp2Test(t *testing.T, impl Dlaqp2er) { _ = "STUB: not implemented"; return }
