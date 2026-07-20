package lp

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

var (
	ErrBland      = errors.New("lp: bland: all replacements are negative or cause ill-conditioned ab")
	ErrInfeasible = errors.New("lp: problem is infeasible")
	ErrLinSolve   = errors.New("lp: linear solve failure")
	ErrUnbounded  = errors.New("lp: problem is unbounded")
	ErrSingular   = errors.New("lp: A is singular")
	ErrZeroColumn = errors.New("lp: A has a column of all zeros")
	ErrZeroRow    = errors.New("lp: A has a row of all zeros")
)

const badShape = "lp: size mismatch"

const (
	initPosTol = 1e-13

	blandNegTol = 1e-14

	rRoundTol = 1e-13

	dRoundTol = 1e-13

	phaseIZeroTol = 1e-12

	blandZeroTol = 1e-12
)

func Simplex(c []float64, A mat.Matrix, b []float64, tol float64, initialBasic []int) (optF float64, optX []float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func simplex(initialBasic []int, c []float64, A mat.Matrix, b []float64, tol float64) (float64, []float64, []int, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil, nil
}

func computeMove(move []float64, minIdx int, A mat.Matrix, ab *mat.Dense, xb []float64, nonBasicIdx []int) error {
	_ = "STUB: not implemented"
	return nil
}

func replaceBland(A mat.Matrix, ab *mat.Dense, xb []float64, basicIdxs, nonBasicIdx []int, r, move []float64) (replace, minIdx int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func verifyInputs(initialBasic []int, c []float64, A mat.Matrix, b []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func initializeFromBasic(xb []float64, ab *mat.Dense, b []float64) error {
	_ = "STUB: not implemented"
	return nil
}

func extractColumns(dst *mat.Dense, A mat.Matrix, cols []int) { _ = "STUB: not implemented"; return }

func findInitialBasic(A mat.Matrix, b []float64) ([]int, *mat.Dense, []float64, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func findLinearlyIndependent(A mat.Matrix) []int { _ = "STUB: not implemented"; return nil }
