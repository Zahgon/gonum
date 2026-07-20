package samplemv

import (
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

var _ Sampler = MetropolisHastingser{}

type MHProposal interface {
	ConditionalLogProb(x, y []float64) (prob float64)

	ConditionalRand(x, y []float64) []float64
}

type MetropolisHastingser struct {
	Initial  []float64
	Target   distmv.LogProber
	Proposal MHProposal
	Src      rand.Source

	BurnIn int
	Rate   int
}

func (m MetropolisHastingser) Sample(batch *mat.Dense) { _ = "STUB: not implemented"; return }

func metropolisHastings(batch *mat.Dense, initial []float64, target distmv.LogProber, proposal MHProposal, src rand.Source) {
	_ = "STUB: not implemented"
	return
}

type ProposalNormal struct {
	normal *distmv.Normal
}

func NewProposalNormal(sigma *mat.SymDense, src rand.Source) (*ProposalNormal, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *ProposalNormal) ConditionalLogProb(x, y []float64) (prob float64) {
	_ = "STUB: not implemented"
	return 0
}

func (p *ProposalNormal) ConditionalRand(x, y []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}
