package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Complex64Level2 = Implementation{}

func (Implementation) Cgbmv(trans blas.Transpose, m, n, kL, kU int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cgemv(trans blas.Transpose, m, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cgerc(m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cgeru(m, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Chbmv(uplo blas.Uplo, n, k int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Chemv(uplo blas.Uplo, n int, alpha complex64, a []complex64, lda int, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cher(uplo blas.Uplo, n int, alpha float32, x []complex64, incX int, a []complex64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cher2(uplo blas.Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, a []complex64, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Chpmv(uplo blas.Uplo, n int, alpha complex64, ap []complex64, x []complex64, incX int, beta complex64, y []complex64, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Chpr(uplo blas.Uplo, n int, alpha float32, x []complex64, incX int, ap []complex64) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Chpr2(uplo blas.Uplo, n int, alpha complex64, x []complex64, incX int, y []complex64, incY int, ap []complex64) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctbmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k int, a []complex64, lda int, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctbsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k int, a []complex64, lda int, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctpmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, ap []complex64, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctpsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, ap []complex64, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctrmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctrsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, a []complex64, lda int, x []complex64, incX int) {
	_ = "STUB: not implemented"
	return
}
