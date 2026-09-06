package mat

type HOGSVD struct {
	n int
	v *Dense
	b []Dense

	err error
}

func (gsvd *HOGSVD) succFact() bool { _ = "STUB: not implemented"; return false }

func (gsvd *HOGSVD) Factorize(m ...Matrix) (ok bool) { _ = "STUB: not implemented"; return false }

func (gsvd *HOGSVD) Err() error { _ = "STUB: not implemented"; return nil }

func (gsvd *HOGSVD) Len() int { _ = "STUB: not implemented"; return 0 }

func (gsvd *HOGSVD) UTo(dst *Dense, n int) { _ = "STUB: not implemented"; return }

func (gsvd *HOGSVD) Values(s []float64, n int) []float64 { _ = "STUB: not implemented"; return nil }

func (gsvd *HOGSVD) VTo(dst *Dense) { _ = "STUB: not implemented"; return }
