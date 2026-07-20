package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Float64Level2 = Implementation{}

func (Implementation) Dger(m, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dgbmv(tA blas.Transpose, m, n, kL, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dgemv(tA blas.Transpose, m, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dtrmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dtrsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dsymv(ul blas.Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dtbmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float64, lda int, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dtpmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dtbsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n, k int, a []float64, lda int, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dsbmv(ul blas.Uplo, n, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dsyr(ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dsyr2(ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dtpsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dspmv(ul blas.Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dspr(ul blas.Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Dspr2(ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, ap []float64) {
	_ = "STUB: not implemented"
	return
}
