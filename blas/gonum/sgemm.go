package gonum

import (
	"gonum.org/v1/gonum/blas"
)

func (Implementation) Sgemm(tA, tB blas.Transpose, m, n, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	_ = "STUB: not implemented"
	return
}

func sgemmParallel(aTrans, bTrans bool, m, n, k int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, alpha float32) {
	_ = "STUB: not implemented"
	return
}

func sgemmSerial(aTrans, bTrans bool, m, n, k int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, alpha float32) {
	_ = "STUB: not implemented"
	return
}

func sgemmSerialNotNot(m, n, k int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, alpha float32) {
	_ = "STUB: not implemented"
	return
}

func sgemmSerialTransNot(m, n, k int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, alpha float32) {
	_ = "STUB: not implemented"
	return
}

func sgemmSerialNotTrans(m, n, k int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, alpha float32) {
	_ = "STUB: not implemented"
	return
}

func sgemmSerialTransTrans(m, n, k int, a []float32, lda int, b []float32, ldb int, c []float32, ldc int, alpha float32) {
	_ = "STUB: not implemented"
	return
}

func sliceView32(a []float32, lda, i, j, r, c int) []float32 { _ = "STUB: not implemented"; return nil }
