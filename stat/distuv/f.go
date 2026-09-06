package distuv

import (
	"math/rand/v2"
)

type F struct {
	D1  float64
	D2  float64
	Src rand.Source
}

func (f F) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f F) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (f F) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (f F) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (f F) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (f F) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f F) Variance() float64 { _ = "STUB: not implemented"; return 0 }
