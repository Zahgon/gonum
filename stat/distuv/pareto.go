package distuv

import (
	"math/rand/v2"
)

type Pareto struct {
	Xm float64

	Alpha float64

	Src rand.Source
}

func (p Pareto) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Entropy() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Quantile(prob float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (p Pareto) Variance() float64 { _ = "STUB: not implemented"; return 0 }
