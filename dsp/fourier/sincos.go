package fourier

type DCT struct {
	work []float64
	ifac [15]int
}

func NewDCT(n int) *DCT { _ = "STUB: not implemented"; return nil }

func (t *DCT) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *DCT) Reset(n int) { _ = "STUB: not implemented"; return }

func (t *DCT) Transform(dst, src []float64) []float64 { _ = "STUB: not implemented"; return nil }

type DST struct {
	work []float64
	ifac [15]int
}

func NewDST(n int) *DST { _ = "STUB: not implemented"; return nil }

func (t *DST) Len() int { _ = "STUB: not implemented"; return 0 }

func (t *DST) Reset(n int) { _ = "STUB: not implemented"; return }

func (t *DST) Transform(dst, src []float64) []float64 { _ = "STUB: not implemented"; return nil }
