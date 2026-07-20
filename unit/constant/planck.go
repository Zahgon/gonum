package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const Planck = planckUnits(6.62607015e-34)

type planckUnits float64

func (cnst planckUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst planckUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
