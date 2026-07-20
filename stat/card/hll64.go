package card

import (
	"hash"
)

type HyperLogLog64 struct {
	p uint8
	m uint64

	hash hash.Hash64

	register []uint8
}

func NewHyperLogLog64(prec int, h hash.Hash64) (*HyperLogLog64, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HyperLogLog64) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (h *HyperLogLog64) Union(a, b *HyperLogLog64) error { _ = "STUB: not implemented"; return nil }

func (h *HyperLogLog64) SetHash(fn hash.Hash64) error { _ = "STUB: not implemented"; return nil }

func (h *HyperLogLog64) Count() float64 { _ = "STUB: not implemented"; return 0 }

func rho64q(x uint64, q uint8) uint8 { _ = "STUB: not implemented"; return 0 }

func (h *HyperLogLog64) Reset() { _ = "STUB: not implemented"; return }

func (h *HyperLogLog64) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *HyperLogLog64) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func hash64For(name string) hash.Hash64 { _ = "STUB: not implemented"; return *new(hash.Hash64) }
