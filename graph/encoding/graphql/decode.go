package graphql

import (
	"encoding/json"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/encoding"
)

func Unmarshal(data []byte, uid string, dst encoding.Builder) error {
	_ = "STUB: not implemented"
	return nil
}

type StringIDSetter interface {
	SetIDFromString(uid string) error
}

type LabelSetter interface {
	SetLabel(string)
}

type generator struct {
	dst encoding.Builder

	uidName string

	nodes map[string]graph.Node
}

func (g *generator) walk(src json.RawMessage, node graph.Node, attr string) error {
	_ = "STUB: not implemented"
	return nil
}
