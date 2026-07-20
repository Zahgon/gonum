package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Complex64Level3 = Implementation{}

func (Implementation) Cgemm(tA, tB blas.Transpose, m, n, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Chemm(side blas.Side, uplo blas.Uplo, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cherk(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha float32, a []complex64, lda int, beta float32, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Cher2k(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta float32, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Csymm(side blas.Side, uplo blas.Uplo, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Csyrk(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex64, a []complex64, lda int, beta complex64, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Csyr2k(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex64, a []complex64, lda int, b []complex64, ldb int, beta complex64, c []complex64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctrmm(side blas.Side, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ctrsm(side blas.Side, uplo blas.Uplo, transA blas.Transpose, diag blas.Diag, m, n int, alpha complex64, a []complex64, lda int, b []complex64, ldb int) {
	_ = "STUB: not implemented"
	return
}
