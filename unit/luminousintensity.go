package unit

import (
	"fmt"
)

type LuminousIntensity float64

const Candela LuminousIntensity = 1

func (j LuminousIntensity) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (j LuminousIntensity) LuminousIntensity() LuminousIntensity {
	_ = "STUB: not implemented"
	return *new(LuminousIntensity)
}

func (j *LuminousIntensity) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (j LuminousIntensity) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
