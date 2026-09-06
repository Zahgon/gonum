package distmat

import (
	"math/rand/v2"
	"sync"

	"gonum.org/v1/gonum/mat"
)

type Wishart struct {
	nu  float64
	src rand.Source

	dim     int
	cholv   mat.Cholesky
	logdetv float64
	upper   mat.TriDense

	once sync.Once
	v    *mat.SymDense
}

func NewWishart(v mat.Symmetric, nu float64, src rand.Source) (*Wishart, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (w *Wishart) MeanSymTo(dst *mat.SymDense) { _ = "STUB: not implemented"; return }

func (w *Wishart) ProbSym(x mat.Symmetric) float64 { _ = "STUB: not implemented"; return 0 }

func (w *Wishart) LogProbSym(x mat.Symmetric) float64 { _ = "STUB: not implemented"; return 0 }

func (w *Wishart) LogProbSymChol(cholX *mat.Cholesky) float64 { _ = "STUB: not implemented"; return 0 }

func (w *Wishart) logProbSymChol(cholX *mat.Cholesky) float64 { _ = "STUB: not implemented"; return 0 }

func (w *Wishart) RandSymTo(dst *mat.SymDense) { _ = "STUB: not implemented"; return }

func (w *Wishart) RandCholTo(dst *mat.Cholesky) { _ = "STUB: not implemented"; return }

func (w *Wishart) setV() { _ = "STUB: not implemented"; return }
