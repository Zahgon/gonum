package testlapack

import (
	"testing"
)

type Drscler interface {
	Drscl(n int, a float64, x []float64, incX int)
}

func DrsclTest(t *testing.T, impl Drscler) { _ = "STUB: not implemented"; return }
