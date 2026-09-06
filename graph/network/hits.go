package network

import (
	"gonum.org/v1/gonum/graph"
)

type HubAuthority struct {
	Hub       float64
	Authority float64
}

func HITS(g graph.Directed, tol float64) map[int64]HubAuthority {
	_ = "STUB: not implemented"
	return nil
}
