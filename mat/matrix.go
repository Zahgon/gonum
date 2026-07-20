package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/lapack"
)

type Matrix interface {
	Dims() (r, c int)

	At(i, j int) float64

	T() Matrix
}

type allMatrix interface {
	Reseter
	IsEmpty() bool
	Zero()
}

type denseMatrix interface {
	DiagView() Diagonal
	Tracer
	Normer
}

var (
	_ Matrix       = Transpose{}
	_ Untransposer = Transpose{}
)

type Transpose struct {
	Matrix Matrix
}

func (t Transpose) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t Transpose) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t Transpose) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (t Transpose) Untranspose() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

type Untransposer interface {
	Untranspose() Matrix
}

type UntransposeBander interface {
	UntransposeBand() Banded
}

type UntransposeTrier interface {
	UntransposeTri() Triangular
}

type UntransposeTriBander interface {
	UntransposeTriBand() TriBanded
}

type Mutable interface {
	Set(i, j int, v float64)

	Matrix
}

type RowViewer interface {
	RowView(i int) Vector
}

type RawRowViewer interface {
	RawRowView(i int) []float64
}

type ColViewer interface {
	ColView(j int) Vector
}

type RawColViewer interface {
	RawColView(j int) []float64
}

type ClonerFrom interface {
	CloneFrom(a Matrix)
}

type Reseter interface {
	Reset()
}

type Copier interface {
	Copy(a Matrix) (r, c int)
}

type Grower interface {
	Caps() (r, c int)
	Grow(r, c int) Matrix
}

type RawMatrixSetter interface {
	SetRawMatrix(a blas64.General)
}

type RawMatrixer interface {
	RawMatrix() blas64.General
}

type RawVectorer interface {
	RawVector() blas64.Vector
}

type NonZeroDoer interface {
	DoNonZero(func(i, j int, v float64))
}

type RowNonZeroDoer interface {
	DoRowNonZero(i int, fn func(i, j int, v float64))
}

type ColNonZeroDoer interface {
	DoColNonZero(j int, fn func(i, j int, v float64))
}

type SolveToer interface {
	SolveTo(dst *Dense, trans bool, b Matrix) error
}

func untranspose(a Matrix) (Matrix, bool) { _ = "STUB: not implemented"; return *new(Matrix), false }

func untransposeExtract(a Matrix) (Matrix, bool) {
	_ = "STUB: not implemented"
	return *new(Matrix), false
}

func Col(dst []float64, j int, a Matrix) []float64 { _ = "STUB: not implemented"; return nil }

func Row(dst []float64, i int, a Matrix) []float64 { _ = "STUB: not implemented"; return nil }

func Cond(a Matrix, norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func Det(a Matrix) float64 { _ = "STUB: not implemented"; return 0 }

func Dot(a, b Vector) float64 { _ = "STUB: not implemented"; return 0 }

func Equal(a, b Matrix) bool { _ = "STUB: not implemented"; return false }

func EqualApprox(a, b Matrix, epsilon float64) bool { _ = "STUB: not implemented"; return false }

func LogDet(a Matrix) (det float64, sign float64) { _ = "STUB: not implemented"; return 0, 0 }

func Max(a Matrix) float64 { _ = "STUB: not implemented"; return 0 }

func Min(a Matrix) float64 { _ = "STUB: not implemented"; return 0 }

type Normer interface {
	Norm(norm float64) float64
}

func Norm(a Matrix, norm float64) float64 { _ = "STUB: not implemented"; return 0 }

func normLapack(norm float64, aTrans bool) lapack.MatrixNorm {
	_ = "STUB: not implemented"
	return *new(lapack.MatrixNorm)
}

func Sum(a Matrix) float64 { _ = "STUB: not implemented"; return 0 }

type Tracer interface {
	Trace() float64
}

func Trace(a Matrix) float64 { _ = "STUB: not implemented"; return 0 }

func use(f []float64, l int) []float64 { _ = "STUB: not implemented"; return nil }

func useZeroed(f []float64, l int) []float64 { _ = "STUB: not implemented"; return nil }

func zero(f []float64) { _ = "STUB: not implemented"; return }

func useInt(i []int, l int) []int { _ = "STUB: not implemented"; return nil }
