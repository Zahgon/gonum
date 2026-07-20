package flow

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/linear"
)

func Intervals(g graph.Directed, eid int64) IntervalGraph {
	_ = "STUB: not implemented"
	return *new(IntervalGraph)
}

type IntervalGraph struct {
	Intervals map[int64]*Interval
	nodes     map[int64]graph.Node
	head      graph.Node
	from      map[int64]map[int64]graph.Edge
	to        map[int64]map[int64]graph.Edge
}

func (ig *IntervalGraph) Head() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (ig *IntervalGraph) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (ig *IntervalGraph) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (ig *IntervalGraph) From(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (ig *IntervalGraph) To(id int64) graph.Nodes {
	_ = "STUB: not implemented"
	return *new(graph.Nodes)
}

func (ig *IntervalGraph) HasEdgeBetween(xid int64, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (ig *IntervalGraph) Node(id int64) graph.Node {
	_ = "STUB: not implemented"
	return *new(graph.Node)
}

func (ig *IntervalGraph) HasEdgeFromTo(uid, vid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (ig *IntervalGraph) setEdge(e graph.Edge) { _ = "STUB: not implemented"; return }

type Interval struct {
	head  graph.Node
	id    int64
	nodes map[int64]graph.Node
	from  map[int64]map[int64]graph.Edge
	to    map[int64]map[int64]graph.Edge
}

func (i *Interval) ID() int64 { _ = "STUB: not implemented"; return 0 }

func (i *Interval) Head() graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (i *Interval) Nodes() graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (i *Interval) Edge(uid, vid int64) graph.Edge {
	_ = "STUB: not implemented"
	return *new(graph.Edge)
}

func (i *Interval) From(id int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (i *Interval) To(id int64) graph.Nodes { _ = "STUB: not implemented"; return *new(graph.Nodes) }

func (i *Interval) HasEdgeBetween(xid int64, yid int64) bool {
	_ = "STUB: not implemented"
	return false
}

func (i *Interval) Node(id int64) graph.Node { _ = "STUB: not implemented"; return *new(graph.Node) }

func (i *Interval) HasEdgeFromTo(uid, vid int64) bool { _ = "STUB: not implemented"; return false }

func (i *Interval) findInterval(h graph.Node, g graph.Directed) (map[int64]graph.Node, map[int64]*Interval) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dfsPostorder(g graph.Directed, eid int64, ns *linear.NodeStack, visited map[int64]bool) {
	_ = "STUB: not implemented"
	return
}

func linkIntervals(intervals []*Interval, g graph.Directed, node2interval map[int64]*Interval) IntervalGraph {
	_ = "STUB: not implemented"
	return *new(IntervalGraph)
}
