package testgraph

import (
	"math"

	"gonum.org/v1/gonum/graph"
)

type node int64

func (n node) ID() int64 { _ = "STUB: not implemented"; return 0 }

type line struct {
	F, T graph.Node
	UID  int64
	W    float64
}

func (e line) From() graph.Node         { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e line) To() graph.Node           { _ = "STUB: not implemented"; return *new(graph.Node) }
func (e line) ReversedEdge() graph.Edge { _ = "STUB: not implemented"; return *new(graph.Edge) }
func (e line) ID() int64                { _ = "STUB: not implemented"; return 0 }
func (e line) Weight() float64          { _ = "STUB: not implemented"; return 0 }

var testCases = []struct {
	name string

	nodes []graph.Node

	edges []WeightedLine

	nonexist []graph.Node

	self float64

	absent float64
}{
	{
		name:     "empty",
		nonexist: []graph.Node{node(-1), node(0), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - negative",
		nodes:    []graph.Node{node(-1)},
		nonexist: []graph.Node{node(0), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - zero",
		nodes:    []graph.Node{node(0)},
		nonexist: []graph.Node{node(-1), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - positive",
		nodes:    []graph.Node{node(1)},
		nonexist: []graph.Node{node(-1), node(0)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "one - self loop",
		nodes:    []graph.Node{node(0)},
		edges:    []WeightedLine{line{F: node(0), T: node(0), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "two - positive",
		nodes:    []graph.Node{node(1), node(2)},
		edges:    []WeightedLine{line{F: node(1), T: node(2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(0)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - negative",
		nodes:    []graph.Node{node(-1), node(-2)},
		edges:    []WeightedLine{line{F: node(-1), T: node(-2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(0), node(-3)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - zero spanning",
		nodes:    []graph.Node{node(-1), node(1)},
		edges:    []WeightedLine{line{F: node(-1), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(0), node(2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - zero contiguous",
		nodes:    []graph.Node{node(0), node(1)},
		edges:    []WeightedLine{line{F: node(0), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(2)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "three - positive",
		nodes:    []graph.Node{node(1), node(2), node(3)},
		edges:    []WeightedLine{line{F: node(1), T: node(2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(0)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - negative",
		nodes:    []graph.Node{node(-1), node(-2), node(-3)},
		edges:    []WeightedLine{line{F: node(-1), T: node(-2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(0), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - zero spanning",
		nodes:    []graph.Node{node(-1), node(0), node(1)},
		edges:    []WeightedLine{line{F: node(-1), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-2), node(2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - zero contiguous",
		nodes:    []graph.Node{node(0), node(1), node(2)},
		edges:    []WeightedLine{line{F: node(0), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(3)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:  "three in only",
		nodes: []graph.Node{node(0), node(1), node(2), node(3)},
		edges: []WeightedLine{
			line{F: node(1), T: node(0), UID: 0, W: 0.5},
			line{F: node(2), T: node(0), UID: 1, W: 0.5},
			line{F: node(3), T: node(0), UID: 2, W: 0.5},
		},
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:  "three out only",
		nodes: []graph.Node{node(0), node(1), node(2), node(3)},
		edges: []WeightedLine{
			line{F: node(0), T: node(1), UID: 0, W: 0.5},
			line{F: node(0), T: node(2), UID: 1, W: 0.5},
			line{F: node(0), T: node(3), UID: 2, W: 0.5},
		},
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name: "4-clique - single(non-prepared)",
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i + 1; j < 4; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique+ - single(non-prepared)",
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i; j < 4; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique - single(prepared)",
		nodes: func() []graph.Node {
			const n = 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = node(i)
			}
			return nodes
		}(),
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique+ - single(prepared)",
		nodes: func() []graph.Node {
			const n = 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = node(i)
			}
			return nodes
		}(),
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name: "4-clique - double(non-prepared)",
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
					edges = append(edges, line{F: node(j), T: node(i), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique+ - double(non-prepared)",
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
					edges = append(edges, line{F: node(j), T: node(i), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique - double(prepared)",
		nodes: func() []graph.Node {
			const n = 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = node(i)
			}
			return nodes
		}(),
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
					edges = append(edges, line{F: node(j), T: node(i), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique+ - double(prepared)",
		nodes: func() []graph.Node {
			const n = 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = node(i)
			}
			return nodes
		}(),
		edges: func() []WeightedLine {
			const n = 4
			var uid int64
			var edges []WeightedLine
			for i := 0; i < n; i++ {
				for j := i; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
					edges = append(edges, line{F: node(j), T: node(i), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:  "issue 1686",
		nodes: []graph.Node{node(0), node(1), node(2)},
		edges: []WeightedLine{
			line{F: node(0), T: node(1), UID: 0, W: 0.5},
			line{F: node(1), T: node(2), UID: 0, W: 0.5},
		},
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
}
