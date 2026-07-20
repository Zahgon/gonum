package blas64

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/gonum"
)

var blas64 blas.Float64 = gonum.Implementation{}

func Use(b blas.Float64) { _ = "STUB: not implemented"; return }

func Implementation() blas.Float64 { _ = "STUB: not implemented"; return *new(blas.Float64) }

type Vector struct {
	N    int
	Data []float64
	Inc  int
}

type General struct {
	Rows, Cols int
	Data       []float64
	Stride     int
}

type Band struct {
	Rows, Cols int
	KL, KU     int
	Data       []float64
	Stride     int
}

type Triangular struct {
	Uplo   blas.Uplo
	Diag   blas.Diag
	N      int
	Data   []float64
	Stride int
}

type TriangularBand struct {
	Uplo   blas.Uplo
	Diag   blas.Diag
	N, K   int
	Data   []float64
	Stride int
}

type TriangularPacked struct {
	Uplo blas.Uplo
	Diag blas.Diag
	N    int
	Data []float64
}

type Symmetric struct {
	Uplo   blas.Uplo
	N      int
	Data   []float64
	Stride int
}

type SymmetricBand struct {
	Uplo   blas.Uplo
	N, K   int
	Data   []float64
	Stride int
}

type SymmetricPacked struct {
	Uplo blas.Uplo
	N    int
	Data []float64
}

const (
	negInc    = "blas64: negative vector increment"
	badLength = "blas64: vector length mismatch"
)

func Dot(x, y Vector) float64 { _ = "STUB: not implemented"; return 0 }

func Nrm2(x Vector) float64 { _ = "STUB: not implemented"; return 0 }

func Asum(x Vector) float64 { _ = "STUB: not implemented"; return 0 }

func Iamax(x Vector) int { _ = "STUB: not implemented"; return 0 }

func Swap(x, y Vector) { _ = "STUB: not implemented"; return }

func Copy(x, y Vector) { _ = "STUB: not implemented"; return }

func Axpy(alpha float64, x, y Vector) { _ = "STUB: not implemented"; return }

func Rotg(a, b float64) (c, s, r, z float64) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

func Rotmg(d1, d2, b1, b2 float64) (p blas.DrotmParams, rd1, rd2, rb1 float64) {
	_ = "STUB: not implemented"
	return *new(blas.DrotmParams), 0, 0, 0
}

func Rot(x, y Vector, c, s float64) { _ = "STUB: not implemented"; return }

func Rotm(x, y Vector, p blas.DrotmParams) { _ = "STUB: not implemented"; return }

func Scal(alpha float64, x Vector) { _ = "STUB: not implemented"; return }

func Gemv(t blas.Transpose, alpha float64, a General, x Vector, beta float64, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Gbmv(t blas.Transpose, alpha float64, a Band, x Vector, beta float64, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Trmv(t blas.Transpose, a Triangular, x Vector) { _ = "STUB: not implemented"; return }

func Tbmv(t blas.Transpose, a TriangularBand, x Vector) { _ = "STUB: not implemented"; return }

func Tpmv(t blas.Transpose, a TriangularPacked, x Vector) { _ = "STUB: not implemented"; return }

func Trsv(t blas.Transpose, a Triangular, x Vector) { _ = "STUB: not implemented"; return }

func Tbsv(t blas.Transpose, a TriangularBand, x Vector) { _ = "STUB: not implemented"; return }

func Tpsv(t blas.Transpose, a TriangularPacked, x Vector) { _ = "STUB: not implemented"; return }

func Symv(alpha float64, a Symmetric, x Vector, beta float64, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Sbmv(alpha float64, a SymmetricBand, x Vector, beta float64, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Spmv(alpha float64, a SymmetricPacked, x Vector, beta float64, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Ger(alpha float64, x, y Vector, a General) { _ = "STUB: not implemented"; return }

func Syr(alpha float64, x Vector, a Symmetric) { _ = "STUB: not implemented"; return }

func Spr(alpha float64, x Vector, a SymmetricPacked) { _ = "STUB: not implemented"; return }

func Syr2(alpha float64, x, y Vector, a Symmetric) { _ = "STUB: not implemented"; return }

func Spr2(alpha float64, x, y Vector, a SymmetricPacked) { _ = "STUB: not implemented"; return }

func Gemm(tA, tB blas.Transpose, alpha float64, a, b General, beta float64, c General) {
	_ = "STUB: not implemented"
	return
}

func Symm(s blas.Side, alpha float64, a Symmetric, b General, beta float64, c General) {
	_ = "STUB: not implemented"
	return
}

func Syrk(t blas.Transpose, alpha float64, a General, beta float64, c Symmetric) {
	_ = "STUB: not implemented"
	return
}

func Syr2k(t blas.Transpose, alpha float64, a, b General, beta float64, c Symmetric) {
	_ = "STUB: not implemented"
	return
}

func Trmm(s blas.Side, tA blas.Transpose, alpha float64, a Triangular, b General) {
	_ = "STUB: not implemented"
	return
}

func Trsm(s blas.Side, tA blas.Transpose, alpha float64, a Triangular, b General) {
	_ = "STUB: not implemented"
	return
}
