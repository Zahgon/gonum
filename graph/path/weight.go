package path

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
)

type Weighted interface {
	Weight(xid, yid int64) (w float64, ok bool)
}

type Weighting func(xid, yid int64) (w float64, ok bool)

func UniformCost(g traverse.Graph) Weighting { _ = "STUB: not implemented"; return *new(Weighting) }

type Heuristic func(x, y graph.Node) float64

type HeuristicCoster interface {
	HeuristicCost(x, y graph.Node) float64
}
