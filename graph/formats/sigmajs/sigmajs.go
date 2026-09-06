package sigmajs

import (
	"encoding/json"
)

type Graph struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	ID         string
	Attributes map[string]interface{}
}

var (
	_ json.Marshaler   = (*Node)(nil)
	_ json.Unmarshaler = (*Node)(nil)
)

func (n *Node) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *Node) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type Edge struct {
	ID         string
	Source     string
	Target     string
	Attributes map[string]interface{}
}

var (
	_ json.Marshaler   = (*Edge)(nil)
	_ json.Unmarshaler = (*Edge)(nil)
)

func (e *Edge) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *Edge) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
