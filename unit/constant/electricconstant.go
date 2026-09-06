package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const ElectricConstant = electricConstantUnits(8.854187817620389e-12)

type electricConstantUnits float64

func (cnst electricConstantUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst electricConstantUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
