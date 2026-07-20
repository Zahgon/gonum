package samplemv

import (
	"errors"
	"math/rand/v2"

	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat/distmv"
)

const errLengthMismatch = "samplemv: slice length mismatch"

var (
	_ Sampler = LatinHypercube{}
	_ Sampler = (*Rejection)(nil)
	_ Sampler = IID{}

	_ WeightedSampler = SampleUniformWeighted{}
	_ WeightedSampler = Importance{}
)

type Sampler interface {
	Sample(batch *mat.Dense)
}

type WeightedSampler interface {
	SampleWeighted(batch *mat.Dense, weights []float64)
}

type SampleUniformWeighted struct {
	Sampler
}

func (w SampleUniformWeighted) SampleWeighted(batch *mat.Dense, weights []float64) {
	_ = "STUB: not implemented"
	return
}

type LatinHypercube struct {
	Q   distmv.Quantiler
	Src rand.Source
}

func (l LatinHypercube) Sample(batch *mat.Dense) { _ = "STUB: not implemented"; return }

func latinHypercube(batch *mat.Dense, q distmv.Quantiler, src rand.Source) {
	_ = "STUB: not implemented"
	return
}

type Importance struct {
	Target   distmv.LogProber
	Proposal distmv.RandLogProber
}

func (l Importance) SampleWeighted(batch *mat.Dense, weights []float64) {
	_ = "STUB: not implemented"
	return
}

func importance(batch *mat.Dense, weights []float64, target distmv.LogProber, proposal distmv.RandLogProber) {
	_ = "STUB: not implemented"
	return
}

var ErrRejection = errors.New("rejection: acceptance ratio above 1")

type Rejection struct {
	C        float64
	Target   distmv.LogProber
	Proposal distmv.RandLogProber
	Src      rand.Source

	err      error
	proposed int
}

func (r *Rejection) Err() error { _ = "STUB: not implemented"; return nil }

func (r *Rejection) Proposed() int { _ = "STUB: not implemented"; return 0 }

func (r *Rejection) Sample(batch *mat.Dense) { _ = "STUB: not implemented"; return }

func rejection(batch *mat.Dense, target distmv.LogProber, proposal distmv.RandLogProber, c float64, src rand.Source) (nProposed int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false
}

type IID struct {
	Dist distmv.Rander
}

func (iid IID) Sample(batch *mat.Dense) { _ = "STUB: not implemented"; return }
