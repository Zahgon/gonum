package distuv

import (
	"math/rand/v2"
)

type Bernoulli struct {
	P   float64
	Src rand.Source
}

func (b Bernoulli) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (Bernoulli) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Bernoulli) Variance() float64 { _ = "STUB: not implemented"; return 0 }
