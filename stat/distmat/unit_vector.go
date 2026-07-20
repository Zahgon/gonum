package distmat

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distuv"
)

type UnitVector struct {
	norm distuv.Normal
}

func NewUnitVector(src rand.Source) *UnitVector { _ = "STUB: not implemented"; return nil }

func (u *UnitVector) UnitVecTo(dst *mat.VecDense) { _ = "STUB: not implemented"; return }
