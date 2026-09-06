package sampleuv

import "math/rand/v2"

type Weighted struct {
	weights []float64

	heap []float64
	rnd  *rand.Rand
}

func NewWeighted(w []float64, src rand.Source) Weighted {
	_ = "STUB: not implemented"
	return *new(Weighted)
}

func (s Weighted) Len() int { _ = "STUB: not implemented"; return 0 }

func (s Weighted) Take() (idx int, ok bool) { _ = "STUB: not implemented"; return 0, false }

func (s Weighted) Reweight(idx int, w float64) { _ = "STUB: not implemented"; return }

func (s Weighted) ReweightAll(w []float64) { _ = "STUB: not implemented"; return }

func (s Weighted) reset() { _ = "STUB: not implemented"; return }
