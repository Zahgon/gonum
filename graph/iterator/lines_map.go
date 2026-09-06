//go:build !safe
// +build !safe

package iterator

import "gonum.org/v1/gonum/graph"

type Lines struct {
	lines int
	iter  *mapIter
	pos   int
	curr  graph.Line
}

func NewLines(lines map[int64]graph.Line) *Lines { _ = "STUB: not implemented"; return nil }

func (l *Lines) Len() int { _ = "STUB: not implemented"; return 0 }

func (l *Lines) Next() bool { _ = "STUB: not implemented"; return false }

func (l *Lines) Line() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func (l *Lines) Reset() { _ = "STUB: not implemented"; return }

func (l *Lines) LineSlice() []graph.Line { _ = "STUB: not implemented"; return nil }

type WeightedLines struct {
	lines int
	iter  *mapIter
	pos   int
	curr  graph.WeightedLine
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
