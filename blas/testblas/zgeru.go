package testblas

import (
	"testing"
)

type Zgeruer interface {
	Zgeru(m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int)
}

func ZgeruTest(t *testing.T, impl Zgeruer) { _ = "STUB: not implemented"; return }
