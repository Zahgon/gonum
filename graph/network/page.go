package network

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/mat"
)

func PageRank(g graph.Directed, damp, tol float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func PageRankSparse(g graph.Directed, damp, tol float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func edgeWeightedPageRank(g graph.WeightedDirected, damp, tol float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func edgeWeightedPageRankSparse(g graph.WeightedDirected, damp, tol float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func pageRank(g graph.Directed, damp, tol float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

func pageRankSparse(g graph.Directed, damp, tol float64) map[int64]float64 {
	_ = "STUB: not implemented"
	return nil
}

type rowCompressedMatrix []compressedRow

func (m rowCompressedMatrix) addTo(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (m rowCompressedMatrix) mulVecUnitary(dst, src *mat.VecDense) {
	_ = "STUB: not implemented"
	return
}

type compressedRow []sparseElement

func (r *compressedRow) addTo(j int, v float64) { _ = "STUB: not implemented"; return }

func (r compressedRow) dotUnitary(v *mat.VecDense) float64 { _ = "STUB: not implemented"; return 0 }

type sparseElement struct {
	index int
	value float64
}

func onesDotUnitary(alpha float64, v *mat.VecDense) float64 { _ = "STUB: not implemented"; return 0 }

func normDiff(x, y []float64) float64 { _ = "STUB: not implemented"; return 0 }
