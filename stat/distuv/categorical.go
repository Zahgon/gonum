package distuv

import (
	"math/rand/v2"
)

type Categorical struct {
	weights []float64

	heap []float64

	src rand.Source
}

func NewCategorical(w []float64, src rand.Source) Categorical {
	_ = "STUB: not implemented"
	return *new(Categorical)
}

func (c Categorical) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Categorical) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (c Categorical) Len() int { _ = "STUB: not implemented"; return 0 }

func (c Categorical) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (c Categorical) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Categorical) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Categorical) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (c Categorical) Reweight(idx int, w float64) { _ = "STUB: not implemented"; return }

func (c Categorical) ReweightAll(w []float64) { _ = "STUB: not implemented"; return }

func (c Categorical) reset() { _ = "STUB: not implemented"; return }
