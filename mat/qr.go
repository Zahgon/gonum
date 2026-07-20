package mat

import (
	"gonum.org/v1/gonum/lapack"
)

const badQR = "mat: invalid QR factorization"

type QR struct {
	qr   *Dense
	q    *Dense
	tau  []float64
	cond float64
}

func (qr *QR) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (qr *QR) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (qr *QR) qRowTo(i int, dst []float64) { _ = "STUB: not implemented"; return }

func (qr *QR) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (qr *QR) updateCond(norm lapack.MatrixNorm) { _ = "STUB: not implemented"; return }

func (qr *QR) Factorize(a Matrix) { _ = "STUB: not implemented"; return }

func (qr *QR) factorize(a Matrix, norm lapack.MatrixNorm) { _ = "STUB: not implemented"; return }

func (qr *QR) updateQ() { _ = "STUB: not implemented"; return }

func (qr *QR) isValid() bool { _ = "STUB: not implemented"; return false }

func (qr *QR) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (qr *QR) RTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (qr *QR) QTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (qr *QR) SolveTo(dst *Dense, trans bool, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func (qr *QR) SolveVecTo(dst *VecDense, trans bool, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}
