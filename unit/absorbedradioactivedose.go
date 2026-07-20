package unit

import (
	"fmt"
)

type AbsorbedRadioactiveDose float64

const Gray AbsorbedRadioactiveDose = 1

func (a AbsorbedRadioactiveDose) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (a AbsorbedRadioactiveDose) AbsorbedRadioactiveDose() AbsorbedRadioactiveDose {
	_ = "STUB: not implemented"
	return *new(AbsorbedRadioactiveDose)
}

func (a *AbsorbedRadioactiveDose) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (a AbsorbedRadioactiveDose) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
