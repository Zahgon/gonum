package prng

type SplitMix64 struct {
	state uint64
}

func NewSplitMix64(seed uint64) *SplitMix64 { _ = "STUB: not implemented"; return nil }

func (src *SplitMix64) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (src *SplitMix64) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (src *SplitMix64) MarshalBinary() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (src *SplitMix64) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

type Xoshiro256plus struct {
	state [4]uint64
}

func NewXoshiro256plus(seed uint64) *Xoshiro256plus { _ = "STUB: not implemented"; return nil }

func (src *Xoshiro256plus) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (src *Xoshiro256plus) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (src *Xoshiro256plus) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (src *Xoshiro256plus) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type Xoshiro256plusplus struct {
	state [4]uint64
}

func NewXoshiro256plusplus(seed uint64) *Xoshiro256plusplus { _ = "STUB: not implemented"; return nil }

func (src *Xoshiro256plusplus) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (src *Xoshiro256plusplus) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (src *Xoshiro256plusplus) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (src *Xoshiro256plusplus) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type Xoshiro256starstar struct {
	state [4]uint64
}

func NewXoshiro256starstar(seed uint64) *Xoshiro256starstar { _ = "STUB: not implemented"; return nil }

func (src *Xoshiro256starstar) Seed(seed uint64) { _ = "STUB: not implemented"; return }

func (src *Xoshiro256starstar) Uint64() uint64 { _ = "STUB: not implemented"; return 0 }

func (src *Xoshiro256starstar) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (src *Xoshiro256starstar) UnmarshalBinary(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}
