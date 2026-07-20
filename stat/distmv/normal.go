package distmv

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type Normal struct {
	mu []float64

	sigma mat.SymDense

	chol       mat.Cholesky
	logSqrtDet float64
	dim        int

	src rand.Source
	rnd *rand.Rand
}

func NewNormal(mu []float64, sigma mat.Symmetric, src rand.Source) (*Normal, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func NewNormalChol(mu []float64, chol *mat.Cholesky, src rand.Source) *Normal {
	_ = "STUB: not implemented"
	return nil
}

func NewNormalPrecision(mu []float64, prec *mat.SymDense, src rand.Source) (norm *Normal, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *Normal) ConditionNormal(observed []int, values []float64, src rand.Source) (*Normal, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *Normal) CovarianceMatrix(dst *mat.SymDense) { _ = "STUB: not implemented"; return }

func (n *Normal) Dim() int { _ = "STUB: not implemented"; return 0 }

func (n *Normal) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (n *Normal) LogProb(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func NormalLogProb(x, mu []float64, chol *mat.Cholesky) float64 {
	_ = "STUB: not implemented"
	return 0
}

func normalLogProb(x, mu []float64, chol *mat.Cholesky, logSqrtDet float64) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (n *Normal) MarginalNormal(vars []int, src rand.Source) (*Normal, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (n *Normal) MarginalNormalSingle(i int, src rand.Source) distuv.Normal {
	_ = "STUB: not implemented"
	return *new(distuv.Normal)
}

func (n *Normal) Mean(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (n *Normal) Prob(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n *Normal) Quantile(dst, p []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (n *Normal) Rand(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func NormalRand(dst, mean []float64, chol *mat.Cholesky, src rand.Source) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type EigenSym interface {
	mat.Symmetric

	RawValues() []float64

	RawQ() mat.Matrix
}

type PositivePartEigenSym struct {
	ed   *mat.EigenSym
	vals []float64
}

var _ EigenSym = (*PositivePartEigenSym)(nil)
var _ EigenSym = (*mat.EigenSym)(nil)

func NewPositivePartEigenSym(ed *mat.EigenSym) *PositivePartEigenSym {
	_ = "STUB: not implemented"
	return nil
}

func (ed *PositivePartEigenSym) SymmetricDim() int { _ = "STUB: not implemented"; return 0 }

func (ed *PositivePartEigenSym) Dims() (r, c int) { _ = "STUB: not implemented"; return 0, 0 }

func (ed *PositivePartEigenSym) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (ed *PositivePartEigenSym) T() mat.Matrix { _ = "STUB: not implemented"; return *new(mat.Matrix) }

func (ed *PositivePartEigenSym) RawQ() mat.Matrix {
	_ = "STUB: not implemented"
	return *new(mat.Matrix)
}

func (ed *PositivePartEigenSym) RawValues() []float64 { _ = "STUB: not implemented"; return nil }

func NormalRandCov(dst, mean []float64, cov mat.Symmetric, src rand.Source) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func (n *Normal) ScoreInput(dst, x []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (n *Normal) SetMean(mu []float64) { _ = "STUB: not implemented"; return }

func (n *Normal) TransformNormal(dst, x []float64) []float64 { _ = "STUB: not implemented"; return nil }

func transformNormal(dst, normal, mu []float64, chol *mat.Cholesky) []float64 {
	_ = "STUB: not implemented"
	return nil
}
