package distmv

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type StudentsT struct {
	nu float64
	mu []float64

	src rand.Source
	rnd *rand.Rand

	sigma mat.SymDense

	chol       mat.Cholesky
	lower      mat.TriDense
	logSqrtDet float64
	dim        int
}

func NewStudentsT(mu []float64, sigma mat.Symmetric, nu float64, src rand.Source) (dist *StudentsT, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *StudentsT) ConditionStudentsT(observed []int, values []float64, src rand.Source) (dist *StudentsT, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func studentsTConditional(observed []int, values []float64, nu float64, mu []float64, sigma mat.Symmetric) (newNu float64, newMean []float64, newSigma *mat.SymDense) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func findUnob(observed []int, dim int) (unobserved []int) { _ = "STUB: not implemented"; return nil }

func (st *StudentsT) CovarianceMatrix(dst *mat.SymDense) { _ = "STUB: not implemented"; return }

func (s *StudentsT) Dim() int { _ = "STUB: not implemented"; return 0 }

func (s *StudentsT) LogProb(y []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s *StudentsT) MarginalStudentsT(vars []int, src rand.Source) (dist *StudentsT, ok bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (s *StudentsT) MarginalStudentsTSingle(i int, src rand.Source) distuv.StudentsT {
	_ = "STUB: not implemented"
	return *new(distuv.StudentsT)
}

func (s *StudentsT) Mean(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (s *StudentsT) Nu() float64 { _ = "STUB: not implemented"; return 0 }

func (s *StudentsT) Prob(y []float64) float64 { _ = "STUB: not implemented"; return 0 }

func (s *StudentsT) Rand(dst []float64) []float64 { _ = "STUB: not implemented"; return nil }
