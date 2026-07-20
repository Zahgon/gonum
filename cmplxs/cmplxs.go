package cmplxs

const (
	zeroLength   = "cmplxs: zero length slice"
	shortSpan    = "cmplxs: slice length less than 2"
	badLength    = "cmplxs: slice lengths do not match"
	badDstLength = "cmplxs: destination slice length does not match input"
)

func Abs(dst []float64, s []complex128) { _ = "STUB: not implemented"; return }

func Add(dst, s []complex128) { _ = "STUB: not implemented"; return }

func AddTo(dst, s, t []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func AddConst(c complex128, dst []complex128) { _ = "STUB: not implemented"; return }

func AddScaled(dst []complex128, alpha complex128, s []complex128) {
	_ = "STUB: not implemented"
	return
}

func AddScaledTo(dst, y []complex128, alpha complex128, s []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func Count(f func(complex128) bool, s []complex128) int { _ = "STUB: not implemented"; return 0 }

func Complex(dst []complex128, real, imag []float64) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func CumProd(dst, s []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func CumSum(dst, s []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func Distance(s, t []complex128, L float64) float64 { _ = "STUB: not implemented"; return 0 }

func Div(dst, s []complex128) { _ = "STUB: not implemented"; return }

func DivTo(dst, s, t []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func Dot(s1, s2 []complex128) complex128 { _ = "STUB: not implemented"; return 0 }

func Equal(s1, s2 []complex128) bool { _ = "STUB: not implemented"; return false }

func EqualApprox(s1, s2 []complex128, tol float64) bool { _ = "STUB: not implemented"; return false }

func EqualFunc(s1, s2 []complex128, f func(complex128, complex128) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func EqualLengths(slices ...[]complex128) bool { _ = "STUB: not implemented"; return false }

func Find(inds []int, f func(complex128) bool, s []complex128, k int) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func HasNaN(s []complex128) bool { _ = "STUB: not implemented"; return false }

func Imag(dst []float64, src []complex128) []float64 { _ = "STUB: not implemented"; return nil }

func LogSpan(dst []complex128, l, u complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func MaxAbs(s []complex128) complex128 { _ = "STUB: not implemented"; return 0 }

func MaxAbsIdx(s []complex128) int { _ = "STUB: not implemented"; return 0 }

func MinAbs(s []complex128) complex128 { _ = "STUB: not implemented"; return 0 }

func MinAbsIdx(s []complex128) int { _ = "STUB: not implemented"; return 0 }

func Mul(dst, s []complex128) { _ = "STUB: not implemented"; return }

func MulConj(dst, s []complex128) { _ = "STUB: not implemented"; return }

func MulConjTo(dst, s, t []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func MulTo(dst, s, t []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func NearestIdx(s []complex128, v complex128) int { _ = "STUB: not implemented"; return 0 }

func Norm(s []complex128, L float64) float64 { _ = "STUB: not implemented"; return 0 }

func Prod(s []complex128) complex128 { _ = "STUB: not implemented"; return 0 }

func Real(dst []float64, src []complex128) []float64 { _ = "STUB: not implemented"; return nil }

func Reverse(s []complex128) { _ = "STUB: not implemented"; return }

func Same(s, t []complex128) bool { _ = "STUB: not implemented"; return false }

func Scale(c complex128, dst []complex128) { _ = "STUB: not implemented"; return }

func ScaleReal(f float64, dst []complex128) { _ = "STUB: not implemented"; return }

func ScaleRealTo(dst []complex128, f float64, s []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func ScaleTo(dst []complex128, c complex128, s []complex128) []complex128 {
	_ = "STUB: not implemented"
	return nil
}

func Span(dst []complex128, l, u complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func Sub(dst, s []complex128) { _ = "STUB: not implemented"; return }

func SubTo(dst, s, t []complex128) []complex128 { _ = "STUB: not implemented"; return nil }

func Sum(s []complex128) complex128 { _ = "STUB: not implemented"; return 0 }
