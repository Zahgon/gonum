package lapack64

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
	"gonum.org/v1/gonum/lapack/gonum"
)

var lapack64 lapack.Float64 = gonum.Implementation{}

func Use(l lapack.Float64) { _ = "STUB: not implemented"; return }

type Tridiagonal struct {
	N  int
	DL []float64
	D  []float64
	DU []float64
}

func Potrf(a blas64.Symmetric) (t blas64.Triangular, ok bool) {
	_ = "STUB: not implemented"
	return *new(blas64.Triangular), false
}

func Potri(t blas64.Triangular) (a blas64.Symmetric, ok bool) {
	_ = "STUB: not implemented"
	return *new(blas64.Symmetric), false
}

func Potrs(t blas64.Triangular, b blas64.General) { _ = "STUB: not implemented"; return }

func Pbcon(a blas64.SymmetricBand, anorm float64, work []float64, iwork []int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Pbtrf(a blas64.SymmetricBand) (t blas64.TriangularBand, ok bool) {
	_ = "STUB: not implemented"
	return *new(blas64.TriangularBand), false
}

func Pbtrs(t blas64.TriangularBand, b blas64.General) { _ = "STUB: not implemented"; return }

func Pstrf(a blas64.Symmetric, piv []int, tol float64, work []float64) (t blas64.Triangular, rank int, ok bool) {
	_ = "STUB: not implemented"
	return *new(blas64.Triangular), 0, false
}

func Gecon(norm lapack.MatrixNorm, a blas64.General, anorm float64, work []float64, iwork []int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Gels(trans blas.Transpose, a blas64.General, b blas64.General, work []float64, lwork int) bool {
	_ = "STUB: not implemented"
	return false
}

func Geqp3(a blas64.General, jpvt []int, tau, work []float64, lwork int) {
	_ = "STUB: not implemented"
	return
}

func Geqrf(a blas64.General, tau, work []float64, lwork int) { _ = "STUB: not implemented"; return }

func Gelqf(a blas64.General, tau, work []float64, lwork int) { _ = "STUB: not implemented"; return }

func Gesvd(jobU, jobVT lapack.SVDJob, a, u, vt blas64.General, s, work []float64, lwork int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Getrf(a blas64.General, ipiv []int) bool { _ = "STUB: not implemented"; return false }

func Getri(a blas64.General, ipiv []int, work []float64, lwork int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Getrs(trans blas.Transpose, a blas64.General, b blas64.General, ipiv []int) {
	_ = "STUB: not implemented"
	return
}

func Ggsvd3(jobU, jobV, jobQ lapack.GSVDJob, a, b blas64.General, alpha, beta []float64, u, v, q blas64.General, work []float64, lwork int, iwork []int) (k, l int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func Gtsv(trans blas.Transpose, a Tridiagonal, b blas64.General) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Lagtm(trans blas.Transpose, alpha float64, a Tridiagonal, b blas64.General, beta float64, c blas64.General) {
	_ = "STUB: not implemented"
	return
}

func Lange(norm lapack.MatrixNorm, a blas64.General, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Langb(norm lapack.MatrixNorm, a blas64.Band) float64 { _ = "STUB: not implemented"; return 0 }

func Langt(norm lapack.MatrixNorm, a Tridiagonal) float64 { _ = "STUB: not implemented"; return 0 }

func Lansb(norm lapack.MatrixNorm, a blas64.SymmetricBand, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Lansy(norm lapack.MatrixNorm, a blas64.Symmetric, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Lantr(norm lapack.MatrixNorm, a blas64.Triangular, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Lantb(norm lapack.MatrixNorm, a blas64.TriangularBand, work []float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Lapmr(forward bool, x blas64.General, k []int) { _ = "STUB: not implemented"; return }

func Lapmt(forward bool, x blas64.General, k []int) { _ = "STUB: not implemented"; return }

func Orglq(a blas64.General, tau, work []float64, lwork int) { _ = "STUB: not implemented"; return }

func Ormlq(side blas.Side, trans blas.Transpose, a blas64.General, tau []float64, c blas64.General, work []float64, lwork int) {
	_ = "STUB: not implemented"
	return
}

func Orgqr(a blas64.General, tau []float64, work []float64, lwork int) {
	_ = "STUB: not implemented"
	return
}

func Ormqr(side blas.Side, trans blas.Transpose, a blas64.General, tau []float64, c blas64.General, work []float64, lwork int) {
	_ = "STUB: not implemented"
	return
}

func Pocon(a blas64.Symmetric, anorm float64, work []float64, iwork []int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Syev(jobz lapack.EVJob, a blas64.Symmetric, w, work []float64, lwork int) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Tbtrs(trans blas.Transpose, a blas64.TriangularBand, b blas64.General) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Trcon(norm lapack.MatrixNorm, a blas64.Triangular, work []float64, iwork []int) float64 {
	_ = "STUB: not implemented"
	return 0
}

func Trtri(a blas64.Triangular) (ok bool) { _ = "STUB: not implemented"; return false }

func Trtrs(trans blas.Transpose, a blas64.Triangular, b blas64.General) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func Geev(jobvl lapack.LeftEVJob, jobvr lapack.RightEVJob, a blas64.General, wr, wi []float64, vl, vr blas64.General, work []float64, lwork int) (first int) {
	_ = "STUB: not implemented"
	return 0
}
