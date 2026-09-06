package prng

const (
	mt19937_64NN        = 312
	mt19937_64MM        = 156
	mt19937_64MatrixA   = 0xB5026F5AA96619E9
	mt19937_64UpperMask = 0xFFFFFFFF80000000
	mt19937_64LowerMask = 0x7FFFFFFF
)

type MT19937_64 struct {
	mt  [mt19937_64NN]uint64
	mti uint64
}

func NewMT19937_64() *MT19937_64 { _ = "STUB: not implemented"; return nil }

func (src *MT19937_64) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (src *MT19937_64) SeedFromKeys(keys []uint64) { _ = "STUB: not implemented"; return }

func (src *MT19937_64) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (src *MT19937_64) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (src *MT19937_64) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
