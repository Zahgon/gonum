package testblas

import (
	"testing"
)

type Izamaxer interface {
	Izamax(n int, x []complex128, incX int) int
}

func IzamaxTest(t *testing.T, impl Izamaxer) { _ = "STUB: not implemented"; return }
