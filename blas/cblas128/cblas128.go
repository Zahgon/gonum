package cblas128

import (
	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/gonum"
)

var cblas128 blas.Complex128 = gonum.Implementation{}

func Use(b blas.Complex128) { _ = "STUB: not implemented"; return }

func Implementation() blas.Complex128 { _ = "STUB: not implemented"; return *new(blas.Complex128) }

type Vector struct {
	N    int
	Inc  int
	Data []complex128
}

type General struct {
	Rows, Cols int
	Stride     int
	Data       []complex128
}

type Band struct {
	Rows, Cols int
	KL, KU     int
	Stride     int
	Data       []complex128
}

type Triangular struct {
	N      int
	Stride int
	Data   []complex128
	Uplo   blas.Uplo
	Diag   blas.Diag
}

type TriangularBand struct {
	N, K   int
	Stride int
	Data   []complex128
	Uplo   blas.Uplo
	Diag   blas.Diag
}

type TriangularPacked struct {
	N    int
	Data []complex128
	Uplo blas.Uplo
	Diag blas.Diag
}

type Symmetric struct {
	N      int
	Stride int
	Data   []complex128
	Uplo   blas.Uplo
}

type SymmetricBand struct {
	N, K   int
	Stride int
	Data   []complex128
	Uplo   blas.Uplo
}

type SymmetricPacked struct {
	N    int
	Data []complex128
	Uplo blas.Uplo
}

type Hermitian Symmetric

type HermitianBand SymmetricBand

type HermitianPacked SymmetricPacked

const (
	negInc    = "cblas128: negative vector increment"
	badLength = "cblas128: vector length mismatch"
)

func Dotu(x, y Vector) complex128 { _ = "STUB: not implemented"; return 0 }

func Dotc(x, y Vector) complex128 { _ = "STUB: not implemented"; return 0 }

func Nrm2(x Vector) float64 { _ = "STUB: not implemented"; return 0 }

func Asum(x Vector) float64 { _ = "STUB: not implemented"; return 0 }

func Iamax(x Vector) int { _ = "STUB: not implemented"; return 0 }

func Swap(x, y Vector) { _ = "STUB: not implemented"; return }

func Copy(x, y Vector) { _ = "STUB: not implemented"; return }

func Axpy(alpha complex128, x, y Vector) { _ = "STUB: not implemented"; return }

func Scal(alpha complex128, x Vector) { _ = "STUB: not implemented"; return }

func Dscal(alpha float64, x Vector) { _ = "STUB: not implemented"; return }

func Gemv(t blas.Transpose, alpha complex128, a General, x Vector, beta complex128, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Gbmv(t blas.Transpose, alpha complex128, a Band, x Vector, beta complex128, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Trmv(t blas.Transpose, a Triangular, x Vector) { _ = "STUB: not implemented"; return }

func Tbmv(t blas.Transpose, a TriangularBand, x Vector) { _ = "STUB: not implemented"; return }

func Tpmv(t blas.Transpose, a TriangularPacked, x Vector) { _ = "STUB: not implemented"; return }

func Trsv(t blas.Transpose, a Triangular, x Vector) { _ = "STUB: not implemented"; return }

func Tbsv(t blas.Transpose, a TriangularBand, x Vector) { _ = "STUB: not implemented"; return }

func Tpsv(t blas.Transpose, a TriangularPacked, x Vector) { _ = "STUB: not implemented"; return }

func Hemv(alpha complex128, a Hermitian, x Vector, beta complex128, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Hbmv(alpha complex128, a HermitianBand, x Vector, beta complex128, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Hpmv(alpha complex128, a HermitianPacked, x Vector, beta complex128, y Vector) {
	_ = "STUB: not implemented"
	return
}

func Geru(alpha complex128, x, y Vector, a General) { _ = "STUB: not implemented"; return }

func Gerc(alpha complex128, x, y Vector, a General) { _ = "STUB: not implemented"; return }

func Her(alpha float64, x Vector, a Hermitian) { _ = "STUB: not implemented"; return }

func Hpr(alpha float64, x Vector, a HermitianPacked) { _ = "STUB: not implemented"; return }

func Her2(alpha complex128, x, y Vector, a Hermitian) { _ = "STUB: not implemented"; return }

func Hpr2(alpha complex128, x, y Vector, a HermitianPacked) { _ = "STUB: not implemented"; return }

func Gemm(tA, tB blas.Transpose, alpha complex128, a, b General, beta complex128, c General) {
	_ = "STUB: not implemented"
	return
}

func Symm(s blas.Side, alpha complex128, a Symmetric, b General, beta complex128, c General) {
	_ = "STUB: not implemented"
	return
}

func Syrk(t blas.Transpose, alpha complex128, a General, beta complex128, c Symmetric) {
	_ = "STUB: not implemented"
	return
}

func Syr2k(t blas.Transpose, alpha complex128, a, b General, beta complex128, c Symmetric) {
	_ = "STUB: not implemented"
	return
}

func Trmm(s blas.Side, tA blas.Transpose, alpha complex128, a Triangular, b General) {
	_ = "STUB: not implemented"
	return
}

func Trsm(s blas.Side, tA blas.Transpose, alpha complex128, a Triangular, b General) {
	_ = "STUB: not implemented"
	return
}

func Hemm(s blas.Side, alpha complex128, a Hermitian, b General, beta complex128, c General) {
	_ = "STUB: not implemented"
	return
}

func Herk(t blas.Transpose, alpha float64, a General, beta float64, c Hermitian) {
	_ = "STUB: not implemented"
	return
}

func Her2k(t blas.Transpose, alpha complex128, a, b General, beta float64, c Hermitian) {
	_ = "STUB: not implemented"
	return
}
