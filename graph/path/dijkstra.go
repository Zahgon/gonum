package path

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/traverse"
)

func DijkstraFrom(u graph.Node, g traverse.Graph) Shortest {
	_ = "STUB: not implemented"
	return *new(Shortest)
}

func DijkstraFromTo(u, t graph.Node, g traverse.Graph) (path []graph.Node, weight float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

func dijkstraFrom(u, t graph.Node, g traverse.Graph) Shortest {
	_ = "STUB: not implemented"
	return *new(Shortest)
}

func DijkstraAllFrom(u graph.Node, g traverse.Graph) ShortestAlts {
	_ = "STUB: not implemented"
	return *new(ShortestAlts)
}

func DijkstraAllPaths(g graph.Graph) (paths AllShortest) {
	_ = "STUB: not implemented"
	return *new(AllShortest)
}

func dijkstraAllPaths(g graph.Graph, paths AllShortest) { _ = "STUB: not implemented"; return }

type distanceNode struct {
	node graph.Node
	dist float64
}

type priorityQueue []distanceNode

func (q priorityQueue) Len() int            { _ = "STUB: not implemented"; return 0 }
func (q priorityQueue) Less(i, j int) bool  { _ = "STUB: not implemented"; return false }
func (q priorityQueue) Swap(i, j int)       { _ = "STUB: not implemented"; return }
func (q *priorityQueue) Push(n interface{}) { _ = "STUB: not implemented"; return }
func (q *priorityQueue) Pop() interface{}   { _ = "STUB: not implemented"; return nil }
