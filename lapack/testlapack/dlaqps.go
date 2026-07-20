package testlapack

import (
	"testing"
)

type Dlaqpser interface {
	Dlapmter
	Dlaqps(m, n, offset, nb int, a []float64, lda int, jpvt []int, tau, vn1, vn2, auxv, f []float64, ldf int) (kb int)
}

func DlaqpsTest(t *testing.T, impl Dlaqpser) { _ = "STUB: not implemented"; return }
