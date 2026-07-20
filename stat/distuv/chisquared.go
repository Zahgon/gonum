package distuv

import (
	"math/rand/v2"
)

type ChiSquared struct {
	K float64

	Src rand.Source
}

func (c ChiSquared) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c ChiSquared) Variance() float64 { _ = "STUB: not implemented"; return 0 }
