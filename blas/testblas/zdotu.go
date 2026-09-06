package testblas

import (
	"testing"
)

type Zdotuer interface {
	Zdotu(n int, x []complex128, incX int, y []complex128, incY int) complex128
}

func ZdotuTest(t *testing.T, impl Zdotuer) { _ = "STUB: not implemented"; return }
