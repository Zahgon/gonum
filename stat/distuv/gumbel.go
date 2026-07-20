package distuv

import (
	"math/rand/v2"
)

type GumbelRight struct {
	Mu   float64
	Beta float64
	Src  rand.Source
}

func (g GumbelRight) z(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (GumbelRight) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (GumbelRight) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (g GumbelRight) Variance() float64 { _ = "STUB: not implemented"; return 0 }
