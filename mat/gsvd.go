package mat

import (
	"gonum.org/v1/gonum/blas/blas64"
)

type GSVDKind int

const (
	GSVDNone GSVDKind = 0

	GSVDU GSVDKind = 1 << iota

	GSVDV

	GSVDQ

	GSVDAll = GSVDU | GSVDV | GSVDQ
)

type GSVD struct {
	kind GSVDKind

	r, p, c, k, l int
	s1, s2        []float64
	a, b, u, v, q blas64.General

	work  []float64
	iwork []int
}

func (gsvd *GSVD) succFact() bool { _ = "STUB: not implemented"; return false }

func (gsvd *GSVD) Factorize(a, b Matrix, kind GSVDKind) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

func (gsvd *GSVD) Kind() GSVDKind { _ = "STUB: not implemented"; return *new(GSVDKind) }

func (gsvd *GSVD) Rank() (k, l int) { _ = "STUB: not implemented"; return 0, 0 }

func (gsvd *GSVD) GeneralizedValues(v []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (gsvd *GSVD) ValuesA(s []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (gsvd *GSVD) ValuesB(s []float64) []float64 { _ = "STUB: not implemented"; return nil }

func (gsvd *GSVD) ZeroRTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (gsvd *GSVD) SigmaATo(dst *Dense) { _ = "STUB: not implemented"; return }

func (gsvd *GSVD) SigmaBTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (gsvd *GSVD) UTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (gsvd *GSVD) VTo(dst *Dense) { _ = "STUB: not implemented"; return }

func (gsvd *GSVD) QTo(dst *Dense) { _ = "STUB: not implemented"; return }
