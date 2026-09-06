package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const Boltzmann = boltzmannUnits(1.380649e-23)

type boltzmannUnits float64

func (cnst boltzmannUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst boltzmannUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
