package iterator

import "gonum.org/v1/gonum/graph"

type OrderedLines struct {
	idx   int
	lines []graph.Line
}

func NewOrderedLines(lines []graph.Line) *OrderedLines { _ = "STUB: not implemented"; return nil }

func (e *OrderedLines) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *OrderedLines) Next() bool { _ = "STUB: not implemented"; return false }

func (e *OrderedLines) Line() graph.Line { _ = "STUB: not implemented"; return *new(graph.Line) }

func (e *OrderedLines) LineSlice() []graph.Line { _ = "STUB: not implemented"; return nil }

func (e *OrderedLines) Reset() { _ = "STUB: not implemented"; return }

type OrderedWeightedLines struct {
	idx   int
	lines []graph.WeightedLine
}

func NewOrderedWeightedLines(lines []graph.WeightedLine) *OrderedWeightedLines {
	_ = "STUB: not implemented"
	return nil
}

func (e *OrderedWeightedLines) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *OrderedWeightedLines) Next() bool { _ = "STUB: not implemented"; return false }

func (e *OrderedWeightedLines) WeightedLine() graph.WeightedLine {
	_ = "STUB: not implemented"
	return *new(graph.WeightedLine)
}

func (e *OrderedWeightedLines) WeightedLineSlice() []graph.WeightedLine {
	_ = "STUB: not implemented"
	return nil
}

func (e *OrderedWeightedLines) Reset() { _ = "STUB: not implemented"; return }
