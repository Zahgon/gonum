package mat

import "gonum.org/v1/gonum/blas/blas64"

func checkOverlap(a, b blas64.General) bool { _ = "STUB: not implemented"; return false }

func (m *Dense) checkOverlap(a blas64.General) bool { _ = "STUB: not implemented"; return false }

func (m *Dense) checkOverlapMatrix(a Matrix) bool { _ = "STUB: not implemented"; return false }

func (s *SymDense) checkOverlap(a blas64.General) bool { _ = "STUB: not implemented"; return false }

func (s *SymDense) checkOverlapMatrix(a Matrix) bool { _ = "STUB: not implemented"; return false }

func generalFromSymmetric(a blas64.Symmetric) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func (t *TriDense) checkOverlap(a blas64.General) bool { _ = "STUB: not implemented"; return false }

func (t *TriDense) checkOverlapMatrix(a Matrix) bool { _ = "STUB: not implemented"; return false }

func generalFromTriangular(a blas64.Triangular) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func (v *VecDense) checkOverlap(a blas64.Vector) bool { _ = "STUB: not implemented"; return false }

func generalFromVector(a blas64.Vector, r, c int) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}

func (s *SymBandDense) checkOverlap(a blas64.General) bool { _ = "STUB: not implemented"; return false }

//lint:ignore U1000 This will be used when we do shadow checks for banded matrices.
func (s *SymBandDense) checkOverlapMatrix(a Matrix) bool { _ = "STUB: not implemented"; return false }

func generalFromSymmetricBand(a blas64.SymmetricBand) blas64.General {
	_ = "STUB: not implemented"
	return *new(blas64.General)
}
