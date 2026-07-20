package testlapack

import (
	"math/rand/v2"
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

const (
	dlamchE = 0x1p-53
	dlamchB = 2
	dlamchP = dlamchB * dlamchE

	dlamchS = 0x1p-1022

	safmin = dlamchS
	safmax = 1 / safmin
	ulp    = dlamchP
	smlnum = safmin / ulp
	bignum = safmax * ulp
)

type worklen int

const (
	minimumWork worklen = iota
	mediumWork
	optimumWork
)

func (wl worklen) String() string { _ = "STUB: not implemented"; return "" }

func normToString(norm lapack.MatrixNorm) string { _ = "STUB: not implemented"; return "" }

func uploToString(uplo blas.Uplo) string { _ = "STUB: not implemented"; return "" }

func diagToString(diag blas.Diag) string { _ = "STUB: not implemented"; return "" }

func sideToString(side blas.Side) string { _ = "STUB: not implemented"; return "" }

func transToString(trans blas.Transpose) string { _ = "STUB: not implemented"; return "" }

func nanSlice(n int) []float64 { _ = "STUB: not implemented"; return nil }

func randomSlice(n int, rnd *rand.Rand) []float64 { _ = "STUB: not implemented"; return nil }

func nanGeneral(r, c, stride int) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func randomGeneral(r, c, stride int, rnd *rand.Rand) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func randomHessenberg(n, stride int, rnd *rand.Rand) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func randomSchurCanonical(n, stride int, bad bool, rnd *rand.Rand) (t blas64.General, wr, wi []float64) {
	_ = "STUB: not implemented"
	return *new(blas64.General), nil, nil
}

func blockedUpperTriGeneral(r, c, k, l, stride int, kblock bool, rnd *rand.Rand) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func nanTriangular(uplo blas.Uplo, n, stride int) blas64.Triangular {
	_ = "STUB: not implemented"
	return *new(blas64.Triangular)
}

func generalOutsideAllNaN(a blas64.General) bool { _ = "STUB: not implemented"; return false }

func triangularOutsideAllNaN(a blas64.Triangular) bool { _ = "STUB: not implemented"; return false }

func transposeGeneral(a blas64.General) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func columnNorms(m, n int, a []float64, lda int) []float64 { _ = "STUB: not implemented"; return nil }

func extractVMat(m, n int, a []float64, lda int, direct lapack.Direct, store lapack.StoreV) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func constructBidiagonal(uplo blas.Uplo, n int, d, e []float64) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func constructVMat(vMat blas64.General, store lapack.StoreV, direct lapack.Direct) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func constructH(tau []float64, v blas64.General, store lapack.StoreV, direct lapack.Direct) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func constructQ(kind string, m, n int, a []float64, lda int, tau []float64) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func constructQK(kind string, m, n, k int, a []float64, lda int, tau []float64) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func checkBidiagonal(t *testing.T, m, n, nb int, a []float64, lda int, d, e, tauP, tauQ, aCopy []float64) {
	_ = "STUB: not implemented"
	return
}

func constructQPBidiagonal(vect lapack.ApplyOrtho, m, n, nb int, a []float64, lda int, tau []float64) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

//lint:ignore U1000 This is useful for debugging.
func printRowise(a []float64, m, n, lda int, beyond bool) { _ = "STUB: not implemented"; return }

func copyGeneral(dst, src blas64.General) { _ = "STUB: not implemented"; return }

func cloneGeneral(a blas64.General) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func equalGeneral(a, b blas64.General) bool { _ = "STUB: not implemented"; return false }

func equalApproxGeneral(a, b blas64.General, tol float64) bool {
	_ = "STUB: not implemented"
	return false
}

func intsEqual(a, b []int) bool { _ = "STUB: not implemented"; return false }

func randSymBand(uplo blas.Uplo, n, kd, ldab int, rnd *rand.Rand) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func distSymBand(uplo blas.Uplo, n, kd int, a []float64, lda int, b []float64, ldb int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func eye(n, stride int) blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func zeros(m, n, stride int) blas64.General { _ = "STUB: not implemented"; return *new(blas64.General) }

func extract2x2Block(t []float64, ldt int) (a, b, c, d float64) {
	_ = "STUB: not implemented"
	return 0, 0, 0, 0
}

func isSchurCanonical(a, b, c, d float64) bool { _ = "STUB: not implemented"; return false }

func isSchurCanonicalGeneral(t blas64.General) bool { _ = "STUB: not implemented"; return false }

//lint:ignore U1000 This is useful for debugging.
func schurBlockEigenvalues(a, b, c, d float64) (ev1, ev2 complex128) {
	_ = "STUB: not implemented"
	return 0, 0
}

func schurBlockSize(t blas64.General, i int) (size int, first bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func containsComplex(v []complex128, z complex128, tol float64) (found bool, index int) {
	_ = "STUB: not implemented"
	return false, 0
}

func isAllNaN(x []float64) bool { _ = "STUB: not implemented"; return false }

func isUpperHessenberg(h blas64.General) bool { _ = "STUB: not implemented"; return false }

func isUpperTriangular(a blas64.General) bool { _ = "STUB: not implemented"; return false }

func unbalancedSparseGeneral(m, n, stride int, nonzeros int, rnd *rand.Rand) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func rootsOfUnity(n int) []complex128 { _ = "STUB: not implemented"; return nil }

func constructGSVDresults(n, p, m, k, l int, a, b blas64.General, alpha, beta []float64) (zeroR, d1, d2 blas64.General) {
	_ = "STUB: not implemented"
	return *new(blas64.General), *new(blas64.General), *new(blas64.General)
}

func constructGSVPresults(n, p, m, k, l int, a, b blas64.General) (zeroA, zeroB blas64.General) {
	_ = "STUB: not implemented"
	return *new(blas64.General), *new(blas64.General)
}

func distFromIdentity(n int, a []float64, lda int) float64 { _ = "STUB: not implemented"; return 0 }

func sameFloat64(a, b float64) bool { _ = "STUB: not implemented"; return false }

func sameLowerTri(n int, a []float64, lda int, b []float64, ldb int) bool {
	_ = "STUB: not implemented"
	return false
}

func sameUpperTri(n int, a []float64, lda int, b []float64, ldb int) bool {
	_ = "STUB: not implemented"
	return false
}

func svdJobString(job lapack.SVDJob) string { _ = "STUB: not implemented"; return "" }

func residualOrthogonal(q blas64.General, rowwise bool) float64 {
	_ = "STUB: not implemented"
	return 0
}
