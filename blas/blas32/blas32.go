package blas32

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/gonum"
)

var blas32 blas.Float32 = gonum.Implementation{}

func Use(b blas.Float32) { _ = "STUB: not implemented"; return }

func Implementation() blas.Float32 { _ = "STUB: not implemented"; return *new(blas.Float32) }

type Vector struct {
	N    int
	Inc  int
	Data []float32
}

type General struct {
	Rows, Cols int
	Stride     int
	Data       []float32
}

type Band struct {
	Rows, Cols int
	KL, KU     int
	Stride     int
	Data       []float32
}

type Triangular struct {
	N      int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
	Diag   blas.Diag
}

type TriangularBand struct {
	N, K   int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
	Diag   blas.Diag
}

type TriangularPacked struct {
	N    int
	Data []float32
	Uplo blas.Uplo
	Diag blas.Diag
}

type Symmetric struct {
	N      int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
}

type SymmetricBand struct {
	N, K   int
	Stride int
	Data   []float32
	Uplo   blas.Uplo
}

type SymmetricPacked struct {
	N    int
	Data []float32
	Uplo blas.Uplo
}

const (
	negInc    = "blas32: negative vector increment"
	badLength = "blas32: vector length mismatch"
)

func Dot(x, y Vector) float32 { _ = "STUB: not implemented"; return 0 }

func DDot(x, y Vector) float64 { _ = "STUB: not implemented"; return 0 }

func SDDot(alpha float32, x, y Vector) float32 { _ = "STUB: not implemented"; return 0 }

func Nrm2(x Vector) float32 { _ = "STUB: not implemented"; return 0 }

func Asum(x Vector) float32 { _ = "STUB: not implemented"; return 0 }

func Iamax(x Vector) int { _ = "STUB: not implemented"; return 0 }

func Swap(x, y Vector) { _ = "STUB: not implemented"; return }

func Copy(x, y Vector) { _ = "STUB: not implemented"; return }

func Axpy(alpha float32, x, y Vector) { _ = "STUB: not implemented"; return }

func Rotg(a, b float32) (c, s, r, z float32) { _ = "STUB: not implemented"; return 0, 0, 0, 0 }

func Rotmg(d1, d2, b1, b2 float32) (p blas.SrotmParams, rd1, rd2, rb1 float32) {
	_ = "STUB: not implemented"
	return *new(blas.SrotmParams), 0, 0, 0
}

func Rot(n int, x, y Vector, c, s float32) { _ = "STUB: not implemented"; return }

func Rotm(n int, x, y Vector, p blas.SrotmParams) { _ = "STUB: not implemented"; return }

func Scal(alpha float32, x Vector) { _ = "STUB: not implemented"; return }

func Gemv(t blas.Transpose, alpha float32, a General, x Vector, beta float32, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Gbmv(t blas.Transpose, alpha float32, a Band, x Vector, beta float32, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Trmv(t blas.Transpose, a Triangular, x Vector) { _ = "STUB: not implemented"; return }

func Tbmv(t blas.Transpose, a TriangularBand, x Vector) { _ = "STUB: not implemented"; return }

func Tpmv(t blas.Transpose, a TriangularPacked, x Vector) { _ = "STUB: not implemented"; return }

func Trsv(t blas.Transpose, a Triangular, x Vector) { _ = "STUB: not implemented"; return }

func Tbsv(t blas.Transpose, a TriangularBand, x Vector) { _ = "STUB: not implemented"; return }

func Tpsv(t blas.Transpose, a TriangularPacked, x Vector) { _ = "STUB: not implemented"; return }

func Symv(alpha float32, a Symmetric, x Vector, beta float32, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Sbmv(alpha float32, a SymmetricBand, x Vector, beta float32, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Spmv(alpha float32, a SymmetricPacked, x Vector, beta float32, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Ger(alpha float32, x, y Vector, a General) { _ = "STUB: not implemented"; return }

func Syr(alpha float32, x Vector, a Symmetric) { _ = "STUB: not implemented"; return }

func Spr(alpha float32, x Vector, a SymmetricPacked) { _ = "STUB: not implemented"; return }

func Syr2(alpha float32, x, y Vector, a Symmetric) { _ = "STUB: not implemented"; return }

func Spr2(alpha float32, x, y Vector, a SymmetricPacked) { _ = "STUB: not implemented"; return }

func Gemm(tA, tB blas.Transpose, alpha float32, a, b General, beta float32, c General) {
	_ = "STUB: not implemented"
	return
}

func Symm(s blas.Side, alpha float32, a Symmetric, b General, beta float32, c General) {
	_ = "STUB: not implemented"
	return
}

func Syrk(t blas.Transpose, alpha float32, a General, beta float32, c Symmetric) {
	_ = "STUB: not implemented"
	return
}

func Syr2k(t blas.Transpose, alpha float32, a, b General, beta float32, c Symmetric) {
	_ = "STUB: not implemented"
	return
}

func Trmm(s blas.Side, tA blas.Transpose, alpha float32, a Triangular, b General) {
	_ = "STUB: not implemented"
	return
}

func Trsm(s blas.Side, tA blas.Transpose, alpha float32, a Triangular, b General) {
	_ = "STUB: not implemented"
	return
}
