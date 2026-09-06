package testblas

import (
	"testing"
)

type Zdotcer interface {
	Zdotc(n int, x []complex128, incX int, y []complex128, incY int) complex128
}

func ZdotcTest(t *testing.T, impl Zdotcer) { _ = "STUB: not implemented"; return }
