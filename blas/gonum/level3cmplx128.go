package gonum

import (
	"gonum.org/v1/gonum/blas"
)

var _ blas.Complex128Level3 = Implementation{}

func (Implementation) Zgemm(tA, tB blas.Transpose, m, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zhemm(side blas.Side, uplo blas.Uplo, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zherk(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha float64, a []complex128, lda int, beta float64, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zher2k(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta float64, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zsymm(side blas.Side, uplo blas.Uplo, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zsyrk(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex128, a []complex128, lda int, beta complex128, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Zsyr2k(uplo blas.Uplo, trans blas.Transpose, n, k int, alpha complex128, a []complex128, lda int, b []complex128, ldb int, beta complex128, c []complex128, ldc int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztrmm(side blas.Side, uplo blas.Uplo, trans blas.Transpose, diag blas.Diag, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int) {
	_ = "STUB: not implemented"
	return
}

func (Implementation) Ztrsm(side blas.Side, uplo blas.Uplo, transA blas.Transpose, diag blas.Diag, m, n int, alpha complex128, a []complex128, lda int, b []complex128, ldb int) {
	_ = "STUB: not implemented"
	return
}
