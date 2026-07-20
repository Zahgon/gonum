package testblas

import (
	"testing"
)

type Zscaler interface {
	Zscal(n int, alpha complex128, x []complex128, incX int)
}

func ZscalTest(t *testing.T, impl Zscaler) { _ = "STUB: not implemented"; return }
