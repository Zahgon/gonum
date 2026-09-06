package mat

const (
	badTriangle = "mat: invalid triangle"
	badCholesky = "mat: invalid Cholesky factorization"
)

var (
	_ Matrix    = (*Cholesky)(nil)
	_ Symmetric = (*Cholesky)(nil)

	_ Matrix    = (*BandCholesky)(nil)
	_ Symmetric = (*BandCholesky)(nil)
	_ Banded    = (*BandCholesky)(nil)
	_ SymBanded = (*BandCholesky)(nil)

	_ Matrix    = (*PivotedCholesky)(nil)
	_ Symmetric = (*PivotedCholesky)(nil)
)

type Cholesky struct {
	chol *TriDense
	cond float64
}

func (c *Cholesky) updateCond(norm float64) { _ = "STUB: not implemented"; return }

func (ch *Cholesky) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (c *Cholesky) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (c *Cholesky) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (c *Cholesky) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (c *Cholesky) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Cholesky) Factorize(a Symmetric) (ok bool) { _ = "STUB: not implemented"; return false }

func (c *Cholesky) Reset() { _ = "STUB: not implemented"; return }

func (c *Cholesky) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (c *Cholesky) SetFromU(t Triangular) { _ = "STUB: not implemented"; return }

func (c *Cholesky) Clone(chol *Cholesky) { _ = "STUB: not implemented"; return }

func (c *Cholesky) Det() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Cholesky) LogDet() float64 { _ = "STUB: not implemented"; return 0 }

func (c *Cholesky) SolveTo(dst *Dense, b Matrix) error { _ = "STUB: not implemented"; return nil }

func (a *Cholesky) SolveCholTo(dst *Dense, b *Cholesky) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Cholesky) SolveVecTo(dst *VecDense, b Vector) error { _ = "STUB: not implemented"; return nil }

func (c *Cholesky) RawU() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (c *Cholesky) UTo(dst *TriDense) { _ = "STUB: not implemented"; return }

func (c *Cholesky) LTo(dst *TriDense) { _ = "STUB: not implemented"; return }

func (c *Cholesky) ToSym(dst *SymDense) { _ = "STUB: not implemented"; return }

func (c *Cholesky) InverseTo(dst *SymDense) error { _ = "STUB: not implemented"; return nil }

func (c *Cholesky) Scale(f float64, orig *Cholesky) { _ = "STUB: not implemented"; return }

func (c *Cholesky) ExtendVecSym(a *Cholesky, v Vector) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *Cholesky) SymRankOne(orig *Cholesky, alpha float64, x Vector) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *Cholesky) valid() bool { _ = "STUB: not implemented"; return false }

type BandCholesky struct {
	chol *TriBandDense
	cond float64
}

func (ch *BandCholesky) Factorize(a SymBanded) (ok bool) { _ = "STUB: not implemented"; return false }

func (ch *BandCholesky) SolveTo(dst *Dense, b Matrix) error { _ = "STUB: not implemented"; return nil }

func (ch *BandCholesky) SolveVecTo(dst *VecDense, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}

func (ch *BandCholesky) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (ch *BandCholesky) Reset() { _ = "STUB: not implemented"; return }

func (ch *BandCholesky) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (ch *BandCholesky) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (ch *BandCholesky) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (ch *BandCholesky) TBand() Banded { _ = "STUB: not implemented"; return *new(Banded) }

func (ch *BandCholesky) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (ch *BandCholesky) Bandwidth() (kl, ku int) { _ = "STUB: not implemented"; return 0, 0 }

func (ch *BandCholesky) SymBand() (n, k int) { _ = "STUB: not implemented"; return 0, 0 }

func (ch *BandCholesky) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (ch *BandCholesky) Det() float64 { _ = "STUB: not implemented"; return 0 }

func (ch *BandCholesky) LogDet() float64 { _ = "STUB: not implemented"; return 0 }

func (ch *BandCholesky) valid() bool { _ = "STUB: not implemented"; return false }

type PivotedCholesky struct {
	chol          *TriDense
	piv, pivTrans []int
	rank          int

	ok   bool
	cond float64
}

func (c *PivotedCholesky) Factorize(a Symmetric, tol float64) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *PivotedCholesky) reset(n int) { _ = "STUB: not implemented"; return }

func (ch *PivotedCholesky) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (c *PivotedCholesky) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (c *PivotedCholesky) T() Matrix { _ = "STUB: not implemented"; return *new(Matrix) }

func (c *PivotedCholesky) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (c *PivotedCholesky) Rank() int { _ = "STUB: not implemented"; return 0 }

func (c *PivotedCholesky) Cond() float64 { _ = "STUB: not implemented"; return 0 }

func (c *PivotedCholesky) RawU() Triangular { _ = "STUB: not implemented"; return *new(Triangular) }

func (c *PivotedCholesky) UTo(dst *TriDense) { _ = "STUB: not implemented"; return }

func (c *PivotedCholesky) ColumnPivots(dst []int) []int { _ = "STUB: not implemented"; return nil }

func (c *PivotedCholesky) SolveTo(dst *Dense, b Matrix) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *PivotedCholesky) SolveVecTo(dst *VecDense, b Vector) error {
	_ = "STUB: not implemented"
	return nil
}
