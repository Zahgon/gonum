package testlapack

import (
	"testing"
)

type Dlags2er interface {
	Dlags2(upper bool, a1, a2, a3, b1, b2, b3 float64) (csu, snu, csv, snv, csq, snq float64)
}

func Dlags2Test(t *testing.T, impl Dlags2er) { _ = "STUB: not implemented"; return }

func det2x2(a, b, c, d float64) float64 { _ = "STUB: not implemented"; return 0 }
