package testlapack

import (
	"testing"
)

type Dlae2er interface {
	Dlae2(a, b, c float64) (rt1, rt2 float64)
}

func Dlae2Test(t *testing.T, impl Dlae2er) { _ = "STUB: not implemented"; return }
