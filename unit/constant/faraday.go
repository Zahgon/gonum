package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const Faraday = faradayUnits(96485.33212)

type faradayUnits float64

func (cnst faradayUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst faradayUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
