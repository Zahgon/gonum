package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const Avogadro = avogadroUnits(6.02214076e+23)

type avogadroUnits float64

func (cnst avogadroUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst avogadroUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
