package testlapack

import (
	"testing"
)

type Dlarfger interface {
	Dlarfg(n int, alpha float64, x []float64, incX int) (beta, tau float64)
}

func DlarfgTest(t *testing.T, impl Dlarfger) { _ = "STUB: not implemented"; return }
