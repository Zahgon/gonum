package floats

const (
	zeroLength   = "floats: zero length slice"
	shortSpan    = "floats: slice length less than 2"
	badLength    = "floats: slice lengths do not match"
	badDstLength = "floats: destination slice length does not match input"
)

func Add(dst, s []float64) { _ = "STUB: not implemented"; return }

func AddTo(dst, s, t []float64) []float64 { _ = "STUB: not implemented"; return nil }

func AddConst(c float64, dst []float64) { _ = "STUB: not implemented"; return }

func AddScaled(dst []float64, alpha float64, s []float64) { _ = "STUB: not implemented"; return }

func AddScaledTo(dst, y []float64, alpha float64, s []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

type argsort struct {
	s    []float64
	inds []int
}

func (a argsort) Len() int { _ = "STUB: not implemented"; return 0 }

func (a argsort) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (a argsort) Swap(i, j int) { _ = "STUB: not implemented"; return }

func Argsort(dst []float64, inds []int) { _ = "STUB: not implemented"; return }

func ArgsortStable(dst []float64, inds []int) { _ = "STUB: not implemented"; return }

func Count(f func(float64) bool, s []float64) int { _ = "STUB: not implemented"; return 0 }

func CumProd(dst, s []float64) []float64 { _ = "STUB: not implemented"; return nil }

func CumSum(dst, s []float64) []float64 { _ = "STUB: not implemented"; return nil }

func Distance(s, t []float64, L float64) float64 { _ = "STUB: not implemented"; return 0 }

func Div(dst, s []float64) { _ = "STUB: not implemented"; return }

func DivTo(dst, s, t []float64) []float64 { _ = "STUB: not implemented"; return nil }

func Dot(s1, s2 []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Equal(s1, s2 []float64) bool { _ = "STUB: not implemented"; return false }

func EqualApprox(s1, s2 []float64, tol float64) bool { _ = "STUB: not implemented"; return false }

func EqualFunc(s1, s2 []float64, f func(float64, float64) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualLengths(slices ...[]float64) bool { _ = "STUB: not implemented"; return false }

func Find(inds []int, f func(float64) bool, s []float64, k int) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HasNaN(s []float64) bool { _ = "STUB: not implemented"; return false }

func LogSpan(dst []float64, l, u float64) []float64 { _ = "STUB: not implemented"; return nil }

func LogSumExp(s []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Max(s []float64) float64 { _ = "STUB: not implemented"; return 0 }

func MaxIdx(s []float64) int { _ = "STUB: not implemented"; return 0 }

func Min(s []float64) float64 { _ = "STUB: not implemented"; return 0 }

func MinIdx(s []float64) int { _ = "STUB: not implemented"; return 0 }

func Mul(dst, s []float64) { _ = "STUB: not implemented"; return }

func MulTo(dst, s, t []float64) []float64 { _ = "STUB: not implemented"; return nil }

func NearestIdx(s []float64, v float64) int { _ = "STUB: not implemented"; return 0 }

func NearestIdxForSpan(n int, l, u float64, v float64) int { _ = "STUB: not implemented"; return 0 }

func Norm(s []float64, L float64) float64 { _ = "STUB: not implemented"; return 0 }

func Prod(s []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Reverse(s []float64) { _ = "STUB: not implemented"; return }

func Same(s, t []float64) bool { _ = "STUB: not implemented"; return false }

func Scale(c float64, dst []float64) { _ = "STUB: not implemented"; return }

func ScaleTo(dst []float64, c float64, s []float64) []float64 {
	_ = "STUB: not implemented"
	return nil
}

func Span(dst []float64, l, u float64) []float64 { _ = "STUB: not implemented"; return nil }

func Sub(dst, s []float64) { _ = "STUB: not implemented"; return }

func SubTo(dst, s, t []float64) []float64 { _ = "STUB: not implemented"; return nil }

func Sum(s []float64) float64 { _ = "STUB: not implemented"; return 0 }

func Within(s []float64, v float64) int { _ = "STUB: not implemented"; return 0 }

func SumCompensated(s []float64) float64 { _ = "STUB: not implemented"; return 0 }
