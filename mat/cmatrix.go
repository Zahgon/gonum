package mat

import (
	"gonum.org/v1/gonum/blas/cblas128"
)

type CMatrix interface {
	Dims() (r, c int)

	At(i, j int) complex128

	H() CMatrix

	T() CMatrix
}

type RawCMatrixer interface {
	RawCMatrix() cblas128.General
}

var (
	_ CMatrix          = ConjTranspose{}
	_ UnConjTransposer = ConjTranspose{}
)

type ConjTranspose struct {
	CMatrix CMatrix
}

func (t ConjTranspose) At(i, j int) complex128 { _ = "STUB: not implemented"; return 0 }

func (t ConjTranspose) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t ConjTranspose) H() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (t ConjTranspose) T() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (t ConjTranspose) UnConjTranspose() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

type CTranspose struct {
	CMatrix CMatrix
}

func (t CTranspose) At(i, j int) complex128 { _ = "STUB: not implemented"; return 0 }

func (t CTranspose) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (t CTranspose) H() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (t CTranspose) T() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

func (t CTranspose) Untranspose() CMatrix { _ = "STUB: not implemented"; return *new(CMatrix) }

type UnConjTransposer interface {
	UnConjTranspose() CMatrix
}

type CUntransposer interface {
	Untranspose() CMatrix
}

func useC(c []complex128, l int) []complex128 { _ = "STUB: not implemented"; return nil }

func useZeroedC(c []complex128, l int) []complex128 { _ = "STUB: not implemented"; return nil }

func zeroC(c []complex128) { _ = "STUB: not implemented"; return }

func untransposeCmplx(a CMatrix) (u CMatrix, trans, conj bool) {
	_ = "STUB: not implemented"
	return *new(CMatrix), false, false
}

func untransposeExtractCmplx(a CMatrix) (u CMatrix, trans, conj bool) {
	_ = "STUB: not implemented"
	return *new(CMatrix), false, false
}

func CEqual(a, b CMatrix) bool { _ = "STUB: not implemented"; return false }

func CEqualApprox(a, b CMatrix, epsilon float64) bool { _ = "STUB: not implemented"; return false }

func cEqualWithinAbsOrRel(a, b complex128, absTol, relTol float64) bool {
	_ = "STUB: not implemented"
	return false
}

func cEqualWithinAbs(a, b complex128, tol float64) bool { _ = "STUB: not implemented"; return false }

const minNormalFloat64 = 2.2250738585072014e-308

func cEqualWithinRel(a, b complex128, tol float64) bool { _ = "STUB: not implemented"; return false }
