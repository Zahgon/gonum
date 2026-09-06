package distuv

import (
	"math/rand/v2"
)

type Beta struct {
	Alpha float64

	Beta float64

	Src rand.Source
}

func (b Beta) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (b Beta) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (b Beta) Variance() float64 { _ = "STUB: not implemented"; return 0 }
