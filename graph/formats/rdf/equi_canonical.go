package rdf

func Lean(g []*Statement) ([]*Statement, error) { _ = "STUB: not implemented"; return nil, nil }

func removeRedundantBnodes(g []*Statement) []*Statement { _ = "STUB: not implemented"; return nil }

type triple [3]string

func isProperSubset(a, b map[triple]bool) bool { _ = "STUB: not implemented"; return false }

func isEqualEdges(a, b map[triple]bool) bool { _ = "STUB: not implemented"; return false }

func findCandidates(g []*Statement) ([]*Statement, map[string]bool, map[string]map[string]bool, bool) {
	_ = "STUB: not implemented"
	return nil, nil, nil, false
}

func inILF(t string, f map[string]bool) bool { _ = "STUB: not implemented"; return false }

func inIF(t string, f map[string]bool) bool { _ = "STUB: not implemented"; return false }

type dfs struct{}

func lean(strategy *dfs, g []*Statement) []*Statement { _ = "STUB: not implemented"; return nil }

func findCoreEndomorphism(strategy *dfs, g []*Statement, cands map[string]map[string]bool, mu map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func selectivity(s *Statement, cands map[string]map[string]bool, preds map[string]int) int {
	_ = "STUB: not implemented"
	return 0
}

func (st *dfs) evaluate(g, q []*Statement, cands map[string]map[string]bool, mu map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func fixedFrom(cands map[string]map[string]bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func applyMu(g []*Statement, mu map[string]string) []*Statement {
	_ = "STUB: not implemented"
	return nil
}

func (st *dfs) search(g, q []*Statement, cands map[string]map[string]bool, mu map[string]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func isAutomorphism(mu map[string]string) bool { _ = "STUB: not implemented"; return false }

func dom(mu map[string]string) map[string]bool { _ = "STUB: not implemented"; return nil }

func codom(mu map[string]string) map[string]bool { _ = "STUB: not implemented"; return nil }

func isEqualTerms(a, b map[string]bool) bool { _ = "STUB: not implemented"; return false }

func sortByCodom(maps []map[string]string) { _ = "STUB: not implemented"; return }

type orderedByCodom struct {
	maps  []map[string]string
	attrs []attrs
}

type attrs struct {
	blanks map[string]bool
	selfs  int
}

func (m orderedByCodom) Len() int           { _ = "STUB: not implemented"; return 0 }
func (m orderedByCodom) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (m orderedByCodom) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (st *dfs) join(q *Statement, g []*Statement, cands map[string]map[string]bool, m map[string]string) []map[string]string {
	_ = "STUB: not implemented"
	return nil
}
