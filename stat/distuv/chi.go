package distuv

import (
	"math/rand/v2"
)

type Chi struct {
	K float64

	Src rand.Source
}

func (c Chi) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (c Chi) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (c Chi) Variance() float64 { _ = "STUB: not implemented"; return 0 }
