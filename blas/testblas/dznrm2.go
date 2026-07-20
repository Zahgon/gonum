package testblas

import (
	"testing"
)

type Dznrm2er interface {
	Dznrm2(n int, x []complex128, incX int) float64
	Dnrm2er
}

func Dznrm2Test(t *testing.T, impl Dznrm2er) { _ = "STUB: not implemented"; return }
