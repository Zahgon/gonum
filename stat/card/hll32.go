package card

import (
	"hash"
)

type HyperLogLog32 struct {
	p uint8
	m uint32

	hash hash.Hash32

	register []uint8
}

func NewHyperLogLog32(prec int, h hash.Hash32) (*HyperLogLog32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *HyperLogLog32) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (h *HyperLogLog32) Union(a, b *HyperLogLog32) error { _ = "STUB: not implemented"; return nil }

func (h *HyperLogLog32) SetHash(fn hash.Hash32) error { _ = "STUB: not implemented"; return nil }

func (h *HyperLogLog32) Count() float64 { _ = "STUB: not implemented"; return 0 }

func rho32q(x uint32, q uint8) uint8 { _ = "STUB: not implemented"; return 0 }

func (h *HyperLogLog32) Reset() { _ = "STUB: not implemented"; return }

func (h *HyperLogLog32) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *HyperLogLog32) UnmarshalBinary(b []byte) error { _ = "STUB: not implemented"; return nil }

func hash32For(name string) hash.Hash32 { _ = "STUB: not implemented"; return *new(hash.Hash32) }
