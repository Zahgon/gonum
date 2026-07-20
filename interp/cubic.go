package interp

import (
	"gonum.org/v1/gonum/mat"
)

type PiecewiseCubic struct {
	xs []float64

	coeffs mat.Dense

	lastY float64

	lastDyDx float64
}

func (pc *PiecewiseCubic) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (pc *PiecewiseCubic) PredictDerivative(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (pc *PiecewiseCubic) FitWithDerivatives(xs, ys, dydxs []float64) {
	_ = "STUB: not implemented"
	return
}

type AkimaSpline struct {
	cubic PiecewiseCubic
}

func (as *AkimaSpline) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (as *AkimaSpline) PredictDerivative(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (as *AkimaSpline) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }

func akimaSlopes(xs, ys []float64) []float64 { _ = "STUB: not implemented"; return nil }

func akimaWeightedAverage(v1, v2, w1, w2 float64) float64 { _ = "STUB: not implemented"; return 0 }

func akimaWeights(slopes []float64, i int) (float64, float64) {
	_ = "STUB: not implemented"
	return 0, 0
}

type FritschButland struct {
	cubic PiecewiseCubic
}

func (fb *FritschButland) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fb *FritschButland) PredictDerivative(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (fb *FritschButland) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }

func fritschButlandEdgeDerivative(xs, ys, slopes []float64, leftEdge bool) float64 {
	_ = "STUB: not implemented"
	return 0
}

func (pc *PiecewiseCubic) fitWithSecondDerivatives(xs, ys, d2ydx2s []float64) {
	_ = "STUB: not implemented"
	return
}

func makeCubicSplineSecondDerivativeEquations(a mat.MutableBanded, b mat.MutableVector, xs, ys []float64) {
	_ = "STUB: not implemented"
	return
}

type NaturalCubic struct {
	cubic PiecewiseCubic
}

func (nc *NaturalCubic) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (nc *NaturalCubic) PredictDerivative(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (nc *NaturalCubic) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }

type ClampedCubic struct {
	cubic PiecewiseCubic
}

func (cc *ClampedCubic) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (cc *ClampedCubic) PredictDerivative(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (cc *ClampedCubic) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }

type NotAKnotCubic struct {
	cubic PiecewiseCubic
}

func (nak *NotAKnotCubic) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (nak *NotAKnotCubic) PredictDerivative(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (nak *NotAKnotCubic) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }
