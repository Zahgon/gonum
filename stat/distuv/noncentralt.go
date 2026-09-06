package distuv

import (
	"errors"
	"math/rand/v2"
)

type NoncentralT struct {
	Nu float64

	Mu float64

	Src rand.Source
}

func (n NoncentralT) Rand() float64 { _ = "STUB: not implemented"; return 0 }

func (n NoncentralT) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (n NoncentralT) Variance() float64 { _ = "STUB: not implemented"; return 0 }

func (n NoncentralT) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n NoncentralT) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n NoncentralT) CDF(t float64) float64 { _ = "STUB: not implemented"; return 0 }

func (n NoncentralT) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

var (
	errInterval = errors.New("distuv: invalid root bracket")
	errMaxIter  = errors.New("distuv: maximum iterations exceeded")
)

func brent(f func(float64) float64, a, b, tol float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func findBracketMono(f func(float64) float64, guess float64) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func normCDF(x, mu, sigma float64, lowerTail bool) float64 { _ = "STUB: not implemented"; return 0 }

func lgamma(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func gammaADivB(a, b float64) float64 { _ = "STUB: not implemented"; return 0 }
