package testblas

import (
	"testing"
)

type Zdscaler interface {
	Zdscal(n int, alpha float64, x []complex128, incX int)
}

func ZdscalTest(t *testing.T, impl Zdscaler) { _ = "STUB: not implemented"; return }
