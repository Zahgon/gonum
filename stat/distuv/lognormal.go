package distuv

import (
	"math/rand/v2"
)

type LogNormal struct {
	Mu    float64
	Sigma float64
	Src   rand.Source
}

func (l LogNormal) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (LogNormal) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l LogNormal) Variance() float64 { _ = "STUB: not implemented"; return 0 }
