package distmv

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
)

type Dirichlet struct {
	alpha []float64
	dim   int
	src   rand.Source

	lbeta    float64
	sumAlpha float64
}

func NewDirichlet(alpha []float64, src rand.Source) *Dirichlet {
	_ = "STUB: not implemented"
	return nil
}

func (d *Dirichlet) CovarianceMatrix(dst *mat.SymDense) { _ = "STUB: not implemented"; return }

func (d *Dirichlet) genLBeta(alpha []float64) (lbeta, sumAlpha float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (d *Dirichlet) Dim() int { _ = "STUB: not implemented"; return 0 }

func (d *Dirichlet) LogProb(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (d *Dirichlet) Mean(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (d *Dirichlet) Prob(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (d *Dirichlet) Rand(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }
