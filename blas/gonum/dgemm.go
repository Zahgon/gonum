package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (Implementation) Dgemm(tA, tB blas.Transpose, m, n, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	_ = "STUB: not implemented"
	return
}

func dgemmParallel(aTrans, bTrans bool, m, n, k int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, alpha float64) {
	_ = "STUB: not implemented"
	return
}

func dgemmSerial(aTrans, bTrans bool, m, n, k int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, alpha float64) {
	_ = "STUB: not implemented"
	return
}

func dgemmSerialNotNot(m, n, k int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, alpha float64) {
	_ = "STUB: not implemented"
	return
}

func dgemmSerialTransNot(m, n, k int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, alpha float64) {
	_ = "STUB: not implemented"
	return
}

func dgemmSerialNotTrans(m, n, k int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, alpha float64) {
	_ = "STUB: not implemented"
	return
}

func dgemmSerialTransTrans(m, n, k int, a []float64, lda int, b []float64, ldb int, c []float64, ldc int, alpha float64) {
	_ = "STUB: not implemented"
	return
}

func sliceView64(a []float64, lda, i, j, r, c int) []float64 { _ = "STUB: not implemented"; return nil }
