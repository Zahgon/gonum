package distuv

import (
	"math/rand/v2"
)

type Poisson struct {
	Lambda float64

	Src rand.Source
}

func (p Poisson) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (Poisson) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (p Poisson) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Poisson) Variance() float64 { _ = "STUB: not implemented"; return 0 }
