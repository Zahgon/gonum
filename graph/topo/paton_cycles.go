package topo

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/internal/set"
)

func UndirectedCyclesIn(g graph.Undirected) [][]graph.Node { _ = "STUB: not implemented"; return nil }

type sets map[int64]set.Ints[int64]

func (s sets) add(uid, vid int64) { _ = "STUB: not implemented"; return }

func (s sets) has(uid int64) bool { _ = "STUB: not implemented"; return false }
