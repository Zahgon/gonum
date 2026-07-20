package unit

import (
	"fmt"
)

type EquivalentRadioactiveDose float64

const Sievert EquivalentRadioactiveDose = 1

func (a EquivalentRadioactiveDose) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (a EquivalentRadioactiveDose) EquivalentRadioactiveDose() EquivalentRadioactiveDose {
	_ = "STUB: not implemented"
	return *new(EquivalentRadioactiveDose)
}

func (a *EquivalentRadioactiveDose) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (a EquivalentRadioactiveDose) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
