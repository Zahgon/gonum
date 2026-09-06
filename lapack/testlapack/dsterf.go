package testlapack

import (
	"testing"
)

type Dsterfer interface {
	Dsteqrer
	Dlansyer
	Dsterf(n int, d, e []float64) (ok bool)
}

func DsterfTest(t *testing.T, impl Dsterfer) { _ = "STUB: not implemented"; return }
