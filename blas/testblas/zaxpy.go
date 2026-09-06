package testblas

import (
	"testing"
)

type Zaxpyer interface {
	Zaxpy(n int, alpha complex128, x []complex128, incX int, y []complex128, incY int)
}

func ZaxpyTest(t *testing.T, impl Zaxpyer) { _ = "STUB: not implemented"; return }
