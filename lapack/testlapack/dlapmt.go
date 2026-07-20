package testlapack

import (
	"testing"
)

type Dlapmter interface {
	Dlapmt(forward bool, m, n int, x []float64, ldx int, k []int)
}

func DlapmtTest(t *testing.T, impl Dlapmter) { _ = "STUB: not implemented"; return }
