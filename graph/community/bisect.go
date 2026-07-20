package community

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/graph"
)

type Interval struct {
	Low, High float64

	Score float64

	Reduced
}

type Reduced interface {
	Communities() [][]graph.Node
}

func Size(g ReducedGraph) float64 { _ = "STUB: not implemented"; return 0 }

func Weight(g ReducedGraph) float64 { _ = "STUB: not implemented"; return 0 }

func ModularScore(g graph.Graph, score func(ReducedGraph) float64, effort int, src rand.Source) func(float64) (float64, Reduced) {
	_ = "STUB: not implemented"
	return nil
}

func LeidenScore(g graph.Graph, score func(ReducedGraph) float64, effort int, src rand.Source) func(float64) (float64, Reduced) {
	_ = "STUB: not implemented"
	return nil
}

func reducedScore(g graph.Graph, reducer func(graph.Graph, float64, rand.Source) ReducedGraph, score func(ReducedGraph) float64, effort int, src rand.Source) func(float64) (float64, Reduced) {
	_ = "STUB: not implemented"
	return nil
}

func SizeMultiplex(g ReducedMultiplex) float64 { _ = "STUB: not implemented"; return 0 }

func WeightMultiplex(g ReducedMultiplex) float64 { _ = "STUB: not implemented"; return 0 }

func ModularMultiplexScore(g Multiplex, weights []float64, all bool, score func(ReducedMultiplex) float64, effort int, src rand.Source) func(float64) (float64, Reduced) {
	_ = "STUB: not implemented"
	return nil
}

func Profile(fn func(float64) (float64, Reduced), log bool, grain, low, high float64) (profile []Interval, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type nonDecreasing int

func (n nonDecreasing) Error() string { _ = "STUB: not implemented"; return "" }

func bisect(fn func(float64) (float64, Reduced), log bool, grain, low, scoreLow, high, scoreHigh float64, comm Reduced) []Interval {
	_ = "STUB: not implemented"
	return nil
}

func fixUp(profile []Interval) []Interval { _ = "STUB: not implemented"; return nil }

func tooSmall(low, high, grain float64, log bool) bool { _ = "STUB: not implemented"; return false }
