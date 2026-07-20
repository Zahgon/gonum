package testlapack

import (
	"testing"
)

type Dlartger interface {
	Dlartg(f, g float64) (cs, sn, r float64)
}

func DlartgTest(t *testing.T, impl Dlartger) { _ = "STUB: not implemented"; return }
