package functions

import (
	"gonum.org/v1/gonum/mat"
)

type Beale struct{}

func (Beale) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (Beale) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Beale) Hess(dst *mat.SymDense, x []float64) { _ = "STUB: not implemented"; return }

func (Beale) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BiggsEXP2 struct{}

func (BiggsEXP2) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (BiggsEXP2) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BiggsEXP2) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BiggsEXP3 struct{}

func (BiggsEXP3) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (BiggsEXP3) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BiggsEXP3) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BiggsEXP4 struct{}

func (BiggsEXP4) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (BiggsEXP4) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BiggsEXP4) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BiggsEXP5 struct{}

func (BiggsEXP5) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (BiggsEXP5) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BiggsEXP5) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BiggsEXP6 struct{}

func (BiggsEXP6) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (BiggsEXP6) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BiggsEXP6) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Box3D struct{}

func (Box3D) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (Box3D) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Box3D) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BraninHoo struct{}

func (BraninHoo) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (BraninHoo) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BrownBadlyScaled struct{}

func (BrownBadlyScaled) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (BrownBadlyScaled) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BrownBadlyScaled) Hess(dst *mat.SymDense, x []float64) { _ = "STUB: not implemented"; return }

func (BrownBadlyScaled) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type BrownAndDennis struct{}

func (BrownAndDennis) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (BrownAndDennis) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (BrownAndDennis) Hess(dst *mat.SymDense, x []float64) { _ = "STUB: not implemented"; return }

func (BrownAndDennis) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type ExtendedPowellSingular struct{}

func (ExtendedPowellSingular) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (ExtendedPowellSingular) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (ExtendedPowellSingular) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type ExtendedRosenbrock struct{}

func (ExtendedRosenbrock) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (ExtendedRosenbrock) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (ExtendedRosenbrock) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Gaussian struct{}

func (Gaussian) y(i int) (yi float64) { _ = "STUB: not implemented"; return 0 }

func (g Gaussian) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (g Gaussian) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Gaussian) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type GulfResearchAndDevelopment struct{}

func (GulfResearchAndDevelopment) Func(x []float64) (sum float64) {
	_ = "STUB: not implemented"
	return 0
}

func (GulfResearchAndDevelopment) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (GulfResearchAndDevelopment) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type HelicalValley struct{}

func (HelicalValley) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (HelicalValley) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (HelicalValley) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Linear struct{}

func (Linear) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (Linear) Grad(grad, x []float64) []float64 { _ = "STUB: not implemented"; return nil }

type PenaltyI struct{}

func (PenaltyI) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (PenaltyI) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (PenaltyI) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type PenaltyII struct{}

func (PenaltyII) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (PenaltyII) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (PenaltyII) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type PowellBadlyScaled struct{}

func (PowellBadlyScaled) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (PowellBadlyScaled) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (PowellBadlyScaled) Hess(dst *mat.SymDense, x []float64) { _ = "STUB: not implemented"; return }

func (PowellBadlyScaled) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Sphere struct{}

func (Sphere) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (Sphere) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Sphere) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Trigonometric struct{}

func (Trigonometric) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (Trigonometric) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Trigonometric) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type VariablyDimensioned struct{}

func (VariablyDimensioned) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (VariablyDimensioned) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (VariablyDimensioned) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Watson struct{}

func (Watson) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (Watson) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Watson) Hess(dst *mat.SymDense, x []float64) { _ = "STUB: not implemented"; return }

func (Watson) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type Wood struct{}

func (Wood) Func(x []float64) (sum float64) { _ = "STUB: not implemented"; return 0 }

func (Wood) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

func (Wood) Hess(dst *mat.SymDense, x []float64) { _ = "STUB: not implemented"; return }

func (Wood) Minima() []Minimum { _ = "STUB: not implemented"; return nil }

type ConcaveRight struct{}

func (ConcaveRight) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (ConcaveRight) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

type ConcaveLeft struct{}

func (ConcaveLeft) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (ConcaveLeft) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

type Plassmann struct {
	L    float64
	Beta float64
}

func (f Plassmann) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f Plassmann) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }

type YanaiOzawaKaneko struct {
	Beta1 float64
	Beta2 float64
}

func (f YanaiOzawaKaneko) Func(x []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (f YanaiOzawaKaneko) Grad(grad, x []float64) { _ = "STUB: not implemented"; return }
