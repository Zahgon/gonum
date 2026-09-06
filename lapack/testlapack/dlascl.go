package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlascler interface {
	Dlascl(kind lapack.MatrixType, kl, ku int, cfrom, cto float64, m, n int, a []float64, lda int)
}

func DlasclTest(t *testing.T, impl Dlascler) { _ = "STUB: not implemented"; return }
