package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Complex128Level2 = Implementation{}

func (Implementation) Zgbmv(trans blas.Transpose, m, n, kL, kU int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zgemv(trans blas.Transpose, m, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zgerc(m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zgeru(m, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zhbmv(uplo blas.Uplo, n, k int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zhemv(uplo blas.Uplo, n int, alpha complex128, a []complex128, lda int, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zher(uplo blas.Uplo, n int, alpha float64, x []complex128, incX int, a []complex128, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zher2(uplo blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, a []complex128, lda int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zhpmv(uplo blas.Uplo, n int, alpha complex128, ap []complex128, x []complex128, incX int, beta complex128, y []complex128, incY int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zhpr(uplo blas.Uplo, n int, alpha float64, x []complex128, incX int, ap []complex128) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zhpr2(uplo blas.Uplo, n int, alpha complex128, x []complex128, incX int, y []complex128, incY int, ap []complex128) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztbmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k int, a []complex128, lda int, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztbsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n, k int, a []complex128, lda int, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztpmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, ap []complex128, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztpsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, ap []complex128, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztrmv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztrsv(uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, n int, a []complex128, lda int, x []complex128, incX int) {
	_ = "STUB: not implemented"
	return
}
