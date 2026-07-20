package distmv

const (
	badQuantile      = "distmv: quantile not between 0 and 1"
	badOutputLen     = "distmv: output slice is not nil or the correct length"
	badInputLength   = "distmv: input slice length mismatch"
	badSizeMismatch  = "distmv: size mismatch"
	badZeroDimension = "distmv: zero dimensional input"
	nonPosDimension  = "distmv: non-positive dimension input"
)

const logTwoPi = 1.8378770664093454835606594728112352797227949472755668

func reuseAs(dst []float64, n int) []float64 { _ = "STUB: not implemented"; return nil }
