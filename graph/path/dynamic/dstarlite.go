package dynamic

import (
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/path"
)

type DStarLite struct {
	s, t *dStarLiteNode
	last *dStarLiteNode

	model       WorldModel
	queue       dStarLiteQueue
	keyModifier float64

	weight    path.Weighting
	heuristic path.Heuristic
}

type WorldModel interface {
	graph.WeightedBuilder
	graph.WeightedDirected
}

func NewDStarLite(s, t graph.Node, g graph.Graph, h path.Heuristic, m WorldModel) *DStarLite {
	_ = "STUB: not implemented"
	return nil
}

func edgeWeight(weight path.Weighting, uid, vid int64) float64 { _ = "STUB: not implemented"; return 0 }

func (d *DStarLite) keyFor(s *dStarLiteNode) key { _ = "STUB: not implemented"; return *new(key) }

func (d *DStarLite) update(u *dStarLiteNode) { _ = "STUB: not implemented"; return }

func (d *DStarLite) findShortestPath() { _ = "STUB: not implemented"; return }

func (d *DStarLite) Step() bool { _ = "STUB: not implemented"; return false }

func (d *DStarLite) MoveTo(n graph.Node) { _ = "STUB: not implemented"; return }

func (d *DStarLite) UpdateWorld(changes []graph.Edge) { _ = "STUB: not implemented"; return }

func (d *DStarLite) worldNodeFor(n graph.Node) *dStarLiteNode {
	_ = "STUB: not implemented"
	return nil
}

func (d *DStarLite) Here() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (d *DStarLite) Path() (p []graph.Node, weight float64) {
	_ = "STUB: not implemented"
	return nil, 0
}

type key [2]float64

var badKey = key{math.NaN(), math.NaN()}

func (k key) isBadKey() bool { _ = "STUB: not implemented"; return false }

func (k key) less(other key) bool { _ = "STUB: not implemented"; return false }

type dStarLiteNode struct {
	graph.Node
	key key
	idx int
	rhs float64
	g   float64
}

func newDStarLiteNode(n graph.Node) *dStarLiteNode { _ = "STUB: not implemented"; return nil }

func (q *dStarLiteNode) inQueue() bool { _ = "STUB: not implemented"; return false }

type dStarLiteQueue []*dStarLiteNode

func (q dStarLiteQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (q dStarLiteQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (q dStarLiteQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (q *dStarLiteQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (q *dStarLiteQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (q dStarLiteQueue) top() *dStarLiteNode { _ = "STUB: not implemented"; return nil }

func (q *dStarLiteQueue) insert(u *dStarLiteNode, k key) { _ = "STUB: not implemented"; return }

func (q *dStarLiteQueue) update(n *dStarLiteNode, k key) { _ = "STUB: not implemented"; return }

func (q *dStarLiteQueue) remove(n *dStarLiteNode) { _ = "STUB: not implemented"; return }
