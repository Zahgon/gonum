package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

func DtrmvBenchmark(b *testing.B, dtrmv Dtrmver, n, lda, incX int, ul blas.Uplo, tA blas.Transpose, d blas.Diag) {
	_ = "STUB: not implemented"
	return
}
