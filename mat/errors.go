package mat

import (
	"gonum.org/v1/gonum/lapack"
)

type Condition float64

func (c Condition) Error() string { _ = "STUB: not implemented"; return "" }

const ConditionTolerance = 1e16

const (
	CondNorm = lapack.MaxRowSum

	CondNormTrans = lapack.MaxColumnSum
)

const stackTraceBufferSize = 1 << 20

func Maybe(fn func()) (err error) { _ = "STUB: not implemented"; return nil }

func MaybeFloat(fn func() float64) (f float64, err error) { _ = "STUB: not implemented"; return 0, nil }

func MaybeComplex(fn func() complex128) (f complex128, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

type Error struct{ string }

func (err Error) Error() string { _ = "STUB: not implemented"; return "" }

var (
	ErrNegativeDimension   = Error{"mat: negative dimension"}
	ErrIndexOutOfRange     = Error{"mat: index out of range"}
	ErrReuseNonEmpty       = Error{"mat: reuse of non-empty matrix"}
	ErrRowAccess           = Error{"mat: row index out of range"}
	ErrColAccess           = Error{"mat: column index out of range"}
	ErrVectorAccess        = Error{"mat: vector index out of range"}
	ErrZeroLength          = Error{"mat: zero length in matrix dimension"}
	ErrRowLength           = Error{"mat: row length mismatch"}
	ErrColLength           = Error{"mat: col length mismatch"}
	ErrSquare              = Error{"mat: expect square matrix"}
	ErrNormOrder           = Error{"mat: invalid norm order for matrix"}
	ErrSingular            = Error{"mat: matrix is singular"}
	ErrShape               = Error{"mat: dimension mismatch"}
	ErrIllegalStride       = Error{"mat: illegal stride"}
	ErrPivot               = Error{"mat: malformed pivot list"}
	ErrTriangle            = Error{"mat: triangular storage mismatch"}
	ErrTriangleSet         = Error{"mat: triangular set out of bounds"}
	ErrBandwidth           = Error{"mat: bandwidth out of range"}
	ErrBandSet             = Error{"mat: band set out of bounds"}
	ErrDiagSet             = Error{"mat: diagonal set out of bounds"}
	ErrSliceLengthMismatch = Error{"mat: input slice length mismatch"}
	ErrNotPSD              = Error{"mat: input not positive symmetric definite"}
	ErrFailedEigen         = Error{"mat: eigendecomposition not successful"}
)

type ErrorStack struct {
	Err error

	StackTrace string
}

func (err ErrorStack) Error() string { _ = "STUB: not implemented"; return "" }

const badCap = "mat: bad capacity"
