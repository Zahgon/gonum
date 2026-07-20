package distuv

type Bhattacharyya struct{}

func (Bhattacharyya) DistBeta(l, r Beta) float64 { _ = "STUB: not implemented"; return 0 }

func (Bhattacharyya) DistNormal(l, r Normal) float64 { _ = "STUB: not implemented"; return 0 }

type Hellinger struct{}

func (Hellinger) DistBeta(l, r Beta) float64 { _ = "STUB: not implemented"; return 0 }

func (Hellinger) DistNormal(l, r Normal) float64 { _ = "STUB: not implemented"; return 0 }

type KullbackLeibler struct{}

func (KullbackLeibler) DistBeta(l, r Beta) float64 { _ = "STUB: not implemented"; return 0 }

func (KullbackLeibler) DistNormal(l, r Normal) float64 { _ = "STUB: not implemented"; return 0 }
