package distmv

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/spatial/r1"
)

type Uniform struct {
	bounds []r1.Interval
	dim    int
	rnd    *rand.Rand
}

func NewUniform(bnds []r1.Interval, src rand.Source) *Uniform {
	_ = "STUB: not implemented"
	return nil
}

func NewUnitUniform(dim int, src rand.Source) *Uniform { _ = "STUB: not implemented"; return nil }

func (u *Uniform) Bounds(bounds []r1.Interval) []r1.Interval { _ = "STUB: not implemented"; return nil }

func (u *Uniform) CDF(dst, x []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (u *Uniform) Dim() int { _ = "STUB: not implemented"; return 0 }

func (u *Uniform) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (u *Uniform) LogProb(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u *Uniform) Mean(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (u *Uniform) Prob(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (u *Uniform) Rand(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (u *Uniform) Quantile(dst, p []float64) []float64 { _ = "STUB: not implemented"; return nil }
