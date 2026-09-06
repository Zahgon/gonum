package testlapack

import (
	"testing"
)

type Dlas2er interface {
	Dlas2(f, g, h float64) (min, max float64)
}

func Dlas2Test(t *testing.T, impl Dlas2er) { _ = "STUB: not implemented"; return }
