package testlapack

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/lapack"
)

func dlagtm(trans blas.Transpose, m, n int, alpha float64, dl, d, du []float64, b []float64, ldb int, beta float64, c []float64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func dlangt(norm lapack.MatrixNorm, n int, dl, d, du []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func dlansy(norm lapack.MatrixNorm, uplo blas.Uplo, n int, a []float64, lda int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func dlange(norm lapack.MatrixNorm, m, n int, a []float64, lda int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func dlansb(norm lapack.MatrixNorm, uplo blas.Uplo, n, kd int, ab []float64, ldab int, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func dlantr(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, m, n int, a []float64, lda int, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func dlantb(norm lapack.MatrixNorm, uplo blas.Uplo, diag blas.Diag, n, k int, a []float64, lda int, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func dlanst(norm lapack.MatrixNorm, n int, d, e []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}
