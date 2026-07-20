package mat

import (
	"gonum.org/v1/gonum/lapack"
)

const badLQ = "mat: invalid LQ factorization"

type LQ struct {
	lq   *Dense
	q    *Dense
	tau  []float64
	cond float64
}

func (lq *LQ) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (lq *LQ) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (lq *LQ) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (lq *LQ) updateCond(norm lapack.MatrixNorm) { _ = "STUB: not implemented"; return }

func (lq *LQ) Factorize(a Matrix) { _ = "STUB: not implemented"; return }

func (lq *LQ) factorize(a Matrix, norm lapack.MatrixNorm) { _ = "STUB: not implemented"; return }

func (lq *LQ) updateQ() { _ = "STUB: not implemented"; return }

func (lq *LQ) isValid() bool { _ = "STUB: not implemented"; return false }

func (lq *LQ) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (lq *LQ) LTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (lq *LQ) QTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (lq *LQ) SolveTo(dst *Dense, trans bool, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func (lq *LQ) SolveVecTo(dst *VecDense, trans bool, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}
