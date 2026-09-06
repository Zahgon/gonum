package unit

import (
	"fmt"
)

type MagneticFlux float64

const Weber MagneticFlux = 1

func (m MagneticFlux) Unit() *Unit { _ = "STUB: not implemented"; return nil }

func (m MagneticFlux) MagneticFlux() MagneticFlux {
	_ = "STUB: not implemented"
	return *new(MagneticFlux)
}

func (m *MagneticFlux) From(u Uniter) error { _ = "STUB: not implemented"; return nil }

func (m MagneticFlux) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
