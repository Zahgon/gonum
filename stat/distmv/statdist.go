package distmv

import (
	"gonum.org/v1/gonum/spatial/r1"
)

type Bhattacharyya struct{}

func (Bhattacharyya) DistNormal(l, r *Normal) float64 { _ = "STUB: not implemented"; return 0 }

func (Bhattacharyya) DistUniform(l, r *Uniform) float64 { _ = "STUB: not implemented"; return 0 }

func unifLogVolOverlap(b1, b2 []r1.Interval) float64 { _ = "STUB: not implemented"; return 0 }

type CrossEntropy struct{}

func (CrossEntropy) DistNormal(l, r *Normal) float64 { _ = "STUB: not implemented"; return 0 }

type Hellinger struct{}

func (Hellinger) DistNormal(l, r *Normal) float64 { _ = "STUB: not implemented"; return 0 }

type KullbackLeibler struct{}

func (KullbackLeibler) DistDirichlet(l, r *Dirichlet) float64 { _ = "STUB: not implemented"; return 0 }

func (KullbackLeibler) DistNormal(l, r *Normal) float64 { _ = "STUB: not implemented"; return 0 }

func (KullbackLeibler) DistUniform(l, r *Uniform) float64 { _ = "STUB: not implemented"; return 0 }

type Renyi struct {
	Alpha float64
}

func (renyi Renyi) DistNormal(l, r *Normal) float64 { _ = "STUB: not implemented"; return 0 }

type Wasserstein struct{}

func (Wasserstein) DistNormal(l, r *Normal) float64 { _ = "STUB: not implemented"; return 0 }
