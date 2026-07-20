package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Float32Level3 = Implementation{}

func (Implementation) Strsm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssymm(s blas.Side, ul blas.Uplo, m, n int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssyrk(ul blas.Uplo, tA blas.Transpose, n, k int, alpha float32, a []float32, lda int, beta float32, c []float32, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssyr2k(ul blas.Uplo, tA blas.Transpose, n, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Strmm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	_ = "STUB: not implemented"
	return
}
