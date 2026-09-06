package prng

const (
	mt19937N         = 624
	mt19937M         = 397
	mt19937matrixA   = 0x9908b0df
	mt19937UpperMask = 0x80000000
	mt19937LowerMask = 0x7fffffff
)

type MT19937 struct {
	mt  [mt19937N]uint32
	mti uint32
}

func NewMT19937() *MT19937 { _ = "STUB: not implemented"; return nil }

func (src *MT19937) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (src *MT19937) SeedFromKeys(keys []uint32) { _ = "STUB: not implemented"; return }

func (src *MT19937) Uint32() uint32 { _ = "STUB: not implemented"; return 0 }

func (src *MT19937) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (src *MT19937) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (src *MT19937) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }
