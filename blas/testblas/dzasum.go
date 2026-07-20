package testblas

import (
	"testing"
)

type Dzasumer interface {
	Dzasum(n int, x []complex128, incX int) float64
}

func DzasumTest(t *testing.T, impl Dzasumer) { _ = "STUB: not implemented"; return }
