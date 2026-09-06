package cytoscapejs

import (
	"encoding/json"
)

type GraphElem struct {
	Elements []Element     `json:"elements"`
	Layout   interface{}   `json:"layout,omitempty"`
	Style    []interface{} `json:"style,omitempty"`
}

type Element struct {
	Group            string      `json:"group,omitempty"`
	Data             ElemData    `json:"data"`
	Position         *Position   `json:"position,omitempty"`
	RenderedPosition *Position   `json:"renderedPosition,omitempty"`
	Selected         bool        `json:"selected,omitempty"`
	Selectable       bool        `json:"selectable,omitempty"`
	Locked           bool        `json:"locked,omitempty"`
	Grabbable        bool        `json:"grabbable,omitempty"`
	Classes          string      `json:"classes,omitempty"`
	Scratch          interface{} `json:"scratch,omitempty"`
}

type ElemType int

const (
	InvalidElement ElemType = iota - 1
	NodeElement
	EdgeElement
)

func (e Element) Type() (ElemType, error) { _ = "STUB: not implemented"; return *new(ElemType), nil }

type ElemData struct {
	ID         string
	Source     string
	Target     string
	Parent     string
	Attributes map[string]interface{}
}

var (
	_ json.Marshaler   = (*ElemData)(nil)
	_ json.Unmarshaler = (*ElemData)(nil)
)

func (e *ElemData) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *ElemData) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type GraphNodeEdge struct {
	Elements Elements      `json:"elements"`
	Layout   interface{}   `json:"layout,omitempty"`
	Style    []interface{} `json:"style,omitempty"`
}

type Elements struct {
	Nodes []Node `json:"nodes"`
	Edges []Edge `json:"edges"`
}

type Node struct {
	Data             NodeData    `json:"data"`
	Position         *Position   `json:"position,omitempty"`
	RenderedPosition *Position   `json:"renderedPosition,omitempty"`
	Selected         bool        `json:"selected,omitempty"`
	Selectable       bool        `json:"selectable,omitempty"`
	Locked           bool        `json:"locked,omitempty"`
	Grabbable        bool        `json:"grabbable,omitempty"`
	Classes          string      `json:"classes,omitempty"`
	Scratch          interface{} `json:"scratch,omitempty"`
}

type NodeData struct {
	ID         string
	Parent     string
	Attributes map[string]interface{}
}

var (
	_ json.Marshaler   = (*NodeData)(nil)
	_ json.Unmarshaler = (*NodeData)(nil)
)

func (n *NodeData) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (n *NodeData) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type Edge struct {
	Data       EdgeData    `json:"data"`
	Selected   bool        `json:"selected,omitempty"`
	Selectable bool        `json:"selectable,omitempty"`
	Classes    string      `json:"classes,omitempty"`
	Scratch    interface{} `json:"scratch,omitempty"`
}

type EdgeData struct {
	ID         string
	Source     string
	Target     string
	Attributes map[string]interface{}
}

var (
	_ json.Marshaler   = (*EdgeData)(nil)
	_ json.Unmarshaler = (*EdgeData)(nil)
)

func (e *EdgeData) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EdgeData) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

type Position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
