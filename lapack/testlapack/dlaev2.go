package testlapack

import (
	"testing"
)

type Dlaev2er interface {
	Dlaev2(a, b, c float64) (rt1, rt2, cs1, sn1 float64)
}

func Dlaev2Test(t *testing.T, impl Dlaev2er) { _ = "STUB: not implemented"; return }

func mul2by2(a, b [2][2]float64) [2][2]float64 { _ = "STUB: not implemented"; return [2][2]float64{} }
