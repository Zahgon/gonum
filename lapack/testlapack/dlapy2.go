package testlapack

import (
	"testing"
)

type Dlapy2er interface {
	Dlapy2(float64, float64) float64
}

func Dlapy2Test(t *testing.T, impl Dlapy2er) { _ = "STUB: not implemented"; return }
