package testblas

import (
	"testing"
)

type Zswaper interface {
	Zswap(n int, x []complex128, incX int, y []complex128, incY int)
}

func ZswapTest(t *testing.T, impl Zswaper) { _ = "STUB: not implemented"; return }
