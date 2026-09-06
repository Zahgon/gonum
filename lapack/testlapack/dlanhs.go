package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlanhser interface {
	Dlanhs(norm lapack.MatrixNorm, n int, a []float64, lda int, work []float64) float64
}

func DlanhsTest(t *testing.T, impl Dlanhser) { _ = "STUB: not implemented"; return }
