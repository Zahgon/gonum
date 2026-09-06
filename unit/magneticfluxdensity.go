package unit

import (
	"fmt"
)

type MagneticFluxDensity float64

const Tesla MagneticFluxDensity = 1

func (m MagneticFluxDensity) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (m MagneticFluxDensity) MagneticFluxDensity() MagneticFluxDensity {
	_ = "STUB: not implemented"
	return *new(MagneticFluxDensity)
}

func (m *MagneticFluxDensity) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (m MagneticFluxDensity) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
