package mat

func (m *Dense) Add(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Sub(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) MulElem(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) DivElem(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Inverse(a Matrix) error { _ = "STUB: not implemented"; return nil }

func (m *Dense) Mul(a, b Matrix) { _ = "STUB: not implemented"; return }

func strictCopy(m *Dense, a Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Exp(a Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Pow(a Matrix, n int) { _ = "STUB: not implemented"; return }

func (m *Dense) Kronecker(a, b Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Scale(f float64, a Matrix) { _ = "STUB: not implemented"; return }

func (m *Dense) Apply(fn func(i, j int, v float64) float64, a Matrix) {
	_ = "STUB: not implemented"
	return
}

func (m *Dense) RankOne(a Matrix, alpha float64, x, y Vector) { _ = "STUB: not implemented"; return }

func (m *Dense) Outer(alpha float64, x, y Vector) { _ = "STUB: not implemented"; return }
