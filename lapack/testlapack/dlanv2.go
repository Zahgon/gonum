package testlapack

import (
	"testing"
)

type Dlanv2er interface {
	Dlanv2(a, b, c, d float64) (aa, bb, cc, dd float64, rt1r, rt1i, rt2r, rt2i float64, cs, sn float64)
}

func Dlanv2Test(t *testing.T, impl Dlanv2er) { _ = "STUB: not implemented"; return }

func dlanv2Test(t *testing.T, impl Dlanv2er, a, b, c, d float64) { _ = "STUB: not implemented"; return }
