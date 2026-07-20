package testblas

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

func DgemvBenchmark(b *testing.B, impl Dgemver, tA blas.Transpose, m, n, incX, incY int) {
	_ = "STUB: not implemented"
	return
}

func DgerBenchmark(b *testing.B, impl Dgerer, m, n, incX, incY int) {
	_ = "STUB: not implemented"
	return
}

type Sgerer interface {
	Sger(m, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int)
}

func SgerBenchmark(b *testing.B, blasser Sgerer, m, n, incX, incY int) {
	_ = "STUB: not implemented"
	return
}
