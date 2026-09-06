//go:build safe
// +build safe

package iterator

import (
	"reflect"

	"gonum.org/v1/gonum/graph"
)

type Lines struct {
	iter     reflect.MapIter
	pos, len int
	curr     graph.Line
	value    reflect.Value
	lines    reflect.Value
}

func NewLines(lines map[int64]graph.Line) *Lines { _ = "STUB: not implemented"; return nil }

func (l *Lines) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *Lines) Next() bool { _ = "STUB: not implemented"; return false }

func (l *Lines) Line() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func (l *Lines) Reset() { _ = "STUB: not implemented"; return }

func (l *Lines) LineSlice() []graph.Line { _ = "STUB: not implemented"; return nil }

type WeightedLines struct {
	iter     reflect.MapIter
	pos, len int
	curr     graph.WeightedLine
	value    reflect.Value
	lines    reflect.Value
}

func NewWeightedLines(lines map[int64]graph.WeightedLine) *WeightedLines {
	_ = "STUB: not implemented"
	return nil
}

func (l *WeightedLines) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *WeightedLines) Next() bool { _ = "STUB: not implemented"; return false }

func (l *WeightedLines) WeightedLine() graph.WeightedLine {
	_ = "STUB: not implemented"
	return *new(graph.WeightedLine)
}

func (l *WeightedLines) Reset() { _ = "STUB: not implemented"; return }

func (l *WeightedLines) WeightedLineSlice() []graph.WeightedLine {
	_ = "STUB: not implemented"
	return nil
}
