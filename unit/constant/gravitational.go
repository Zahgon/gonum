package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const Gravitational = gravitationalUnits(6.6743e-11)

type gravitationalUnits float64

func (cnst gravitationalUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst gravitationalUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
