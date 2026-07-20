package distuv

type Logistic struct {
	Mu float64
	S  float64
}

func (l Logistic) CDF(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) ExKurtosis() float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) LogProb(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Mean() float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Mode() float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Median() float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) NumParameters() int { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Prob(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Quantile(p float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Skewness() float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) StdDev() float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Survival(x float64) float64 { _ = "STUB: not implemented"; return 0 }

func (l Logistic) Variance() float64 { _ = "STUB: not implemented"; return 0 }
