package interp

const (
	differentLengths        = "interp: input slices have different lengths"
	tooFewPoints            = "interp: too few points for interpolation"
	xsNotStrictlyIncreasing = "interp: xs values not strictly increasing"
)

type Predictor interface {
	Predict(x float64) float64
}

type Fitter interface {
	Fit(xs, ys []float64) error
}

type FittablePredictor interface {
	Fitter
	Predictor
}

type DerivativePredictor interface {
	Predictor

	PredictDerivative(x float64) float64
}

type Constant float64

func (c Constant) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

type Function func(float64) float64

func (fn Function) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

type PiecewiseLinear struct {
	xs []float64

	ys []float64

	slopes []float64
}

func (pl *PiecewiseLinear) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }

func (pl PiecewiseLinear) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

type PiecewiseConstant struct {
	xs []float64

	ys []float64
}

func (pc *PiecewiseConstant) Fit(xs, ys []float64) error { _ = "STUB: not implemented"; return nil }

func (pc PiecewiseConstant) Predict(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func findSegment(xs []float64, x float64) int { _ = "STUB: not implemented"; return 0 }

func calculateSlopes(xs, ys []float64) []float64 { _ = "STUB: not implemented"; return nil }
