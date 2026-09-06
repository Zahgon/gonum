package constant

import (
	"fmt"

	"gonum.org/v1/gonum/unit"
)

const MagneticConstant = magneticConstantUnits(1.2566370621238374e-06)

type magneticConstantUnits float64

func (cnst magneticConstantUnits) Unit() *unit.Unit { _ = "STUB: not implemented"; return nil }

func (cnst magneticConstantUnits) Format(fs fmt.State, c rune) { _ = "STUB: not implemented"; return }
