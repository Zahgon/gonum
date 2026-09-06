package testlapack

import (
	"testing"
)

type Dlasv2er interface {
	Dlasv2(f, g, h float64) (ssmin, ssmax, snr, csr, snl, csl float64)
}

func Dlasv2Test(t *testing.T, impl Dlasv2er) { _ = "STUB: not implemented"; return }
