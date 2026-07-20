package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

type Dlansyer interface {
	Dlanger
	Dlansy(norm lapack.MatrixNorm, uplo blas.Uplo, n int, a []float64, lda int, work []float64) float64
}

func DlansyTest(t *testing.T, impl Dlansyer) { _ = "STUB: not implemented"; return }
