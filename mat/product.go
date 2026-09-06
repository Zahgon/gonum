package mat

func (m *Dense) Product(factors ...Matrix) { _ = "STUB: not implemented"; return }

const debugProductWalk = false

type multiplier struct {
	factors []Matrix

	dims []int

	table table
}

func newMultiplier(m *Dense, factors []Matrix) *multiplier { _ = "STUB: not implemented"; return nil }

func (p *multiplier) optimize() { _ = "STUB: not implemented"; return }

func (p *multiplier) multiply() *Dense { _ = "STUB: not implemented"; return nil }

func (p *multiplier) multiplySubchain(i, j int) (m Matrix, intermediate bool) {
	_ = "STUB: not implemented"
	return *new(Matrix), false
}

type entry struct {
	k    int
	cost int
}

type table struct {
	n       int
	entries []entry
}

func newTable(n int) table { _ = "STUB: not implemented"; return *new(table) }

func (t table) at(i, j int) entry     { _ = "STUB: not implemented"; return *new(entry) }
func (t table) set(i, j int, e entry) { _ = "STUB: not implemented"; return }

type result bool

func (r result) String() string { _ = "STUB: not implemented"; return "" }
