package distuv

import (
	"math/rand/v2"
)

type Binomial struct {
	N float64

	P float64

	Src rand.Source
}

func (b Binomial) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (Binomial) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (b Binomial) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Binomial) Variance() float64 { _ = "STUB: not implemented"; return 0 }
