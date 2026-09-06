package distmat

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
)

type UniformPermutation struct {
	rnd     *rand.Rand
	indices []int
}

func NewUniformPermutation(src rand.Source) *UniformPermutation {
	_ = "STUB: not implemented"
	return nil
}

func (p *UniformPermutation) PermTo(dst *mat.Dense) { _ = "STUB: not implemented"; return }
