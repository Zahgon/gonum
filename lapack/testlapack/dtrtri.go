package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dtrtrier interface {
	Dtrtri(uplo blas.Uplo, diag blas.Diag, n int, a []float64, lda int) bool
}

func DtrtriTest(t *testing.T, impl Dtrtrier) { _ = "STUB: not implemented"; return }
