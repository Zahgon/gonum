package distuv

type Parameter struct {
	Name  string
	Value float64
}

const (
	badPercentile = "distuv: percentile out of bounds"
	badLength     = "distuv: slice length mismatch"
	badSuffStat   = "distuv: wrong suffStat length"
	errNoSamples  = "distuv: must have at least one sample"
)

const (
	expNegOneHalf   = 0.6065306597126334236037995349911804534419
	eulerMascheroni = 0.5772156649015328606065120900824024310421
	apery           = 1.2020569031595942853997381615114499907649
)
