package testblas

import (
	"testing"
)

type Zcopyer interface {
	Zcopy(n int, x []complex128, incX int, y []complex128, incY int)
}

func ZcopyTest(t *testing.T, impl Zcopyer) { _ = "STUB: not implemented"; return }
