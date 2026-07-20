package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlasrter interface {
	Dlasrt(s lapack.Sort, n int, d []float64)
}

func DlasrtTest(t *testing.T, impl Dlasrter) { _ = "STUB: not implemented"; return }
