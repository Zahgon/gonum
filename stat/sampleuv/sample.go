package sampleuv

import (
	"errors"
	"math/rand/v2"

	"gonum.org/v1/gonum/stat/distuv"
)

const badLengthMismatch = "sample: slice length mismatch"

var (
	_ Sampler = LatinHypercube{}
	_ Sampler = MetropolisHastings{}
	_ Sampler = (*Rejection)(nil)
	_ Sampler = IIDer{}

	_ WeightedSampler = SampleUniformWeighted{}
	_ WeightedSampler = Importance{}
)

type Sampler interface {
	Sample(batch []float64)
}

type WeightedSampler interface {
	SampleWeighted(batch, weights []float64)
}

type SampleUniformWeighted struct {
	Sampler
}

func (w SampleUniformWeighted) SampleWeighted(batch, weights []float64) {
	_ = "STUB: not implemented"
	return
}

type LatinHypercube struct {
	Q   distuv.Quantiler
	Src rand.Source
}

func (l LatinHypercube) Sample(batch []float64) { _ = "STUB: not implemented"; return }

func latinHypercube(batch []float64, q distuv.Quantiler, src rand.Source) {
	_ = "STUB: not implemented"
	return
}

type Importance struct {
	Target   distuv.LogProber
	Proposal distuv.RandLogProber
}

func (l Importance) SampleWeighted(batch, weights []float64) { _ = "STUB: not implemented"; return }

func importance(batch, weights []float64, target distuv.LogProber, proposal distuv.RandLogProber) {
	_ = "STUB: not implemented"
	return
}

var ErrRejection = errors.New("rejection: acceptance ratio above 1")

type Rejection struct {
	C        float64
	Target   distuv.LogProber
	Proposal distuv.RandLogProber
	Src      rand.Source

	err      error
	proposed int
}

func (r *Rejection) Err() error { _ = "STUB: not implemented"; return nil }

func (r *Rejection) Proposed() int { _ = "STUB: not implemented"; return 0 }

func (r *Rejection) Sample(batch []float64) { _ = "STUB: not implemented"; return }

func rejection(batch []float64, target distuv.LogProber, proposal distuv.RandLogProber, c float64, src rand.Source) (nProposed int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type MHProposal interface {
	ConditionalLogProb(x, y float64) (prob float64)

	ConditionalRand(y float64) (x float64)
}

type MetropolisHastings struct {
	Initial  float64
	Target   distuv.LogProber
	Proposal MHProposal
	Src      rand.Source

	BurnIn int
	Rate   int
}

func (m MetropolisHastings) Sample(batch []float64) { _ = "STUB: not implemented"; return }

func metropolisHastings(batch []float64, initial float64, target distuv.LogProber, proposal MHProposal, src rand.Source) {
	_ = "STUB: not implemented"
	return
}

type IIDer struct {
	Dist distuv.Rander
}

func (iid IIDer) Sample(batch []float64) { _ = "STUB: not implemented"; return }
