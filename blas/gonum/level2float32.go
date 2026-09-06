package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Float32Level2 = Implementation{}

func (Implementation) Sger(m, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Sgbmv(tA blas.Transpose, m, n, kL, kU int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Sgemv(tA blas.Transpose, m, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Strmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float32, lda int, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Strsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float32, lda int, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssymv(ul blas.Uplo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Stbmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float32, lda int, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Stpmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float32, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Stbsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float32, lda int, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssbmv(ul blas.Uplo, n, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssyr(ul blas.Uplo, n int, alpha float32, x []float32, incX int, a []float32, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ssyr2(ul blas.Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Stpsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float32, x []float32, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Sspmv(ul blas.Uplo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Sspr(ul blas.Uplo, n int, alpha float32, x []float32, incX int, ap []float32) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Sspr2(ul blas.Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, ap []float32) {
	_ = "STUB: not implemented"
	return
}
