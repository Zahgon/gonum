package mat

import (
	"sync"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/blas/cblas128"
)

func poolFor(size uint) int { _ = "STUB: not implemented"; return 0 }

var (
	poolDense [63]sync.Pool

	poolSymDense [63]sync.Pool

	poolTriDense [63]sync.Pool

	poolVecDense [63]sync.Pool

	poolCDense [63]sync.Pool

	poolFloat64s [63]sync.Pool

	poolInts [63]sync.Pool
)

func init() {
	for i := range poolDense {
		l := 1 << uint(i)

		poolDense[i].New = func() interface{} {
			return &Dense{mat: blas64.General{
				Data: make([]float64, l),
			}}
		}
		poolSymDense[i].New = func() interface{} {
			return &SymDense{mat: blas64.Symmetric{
				Uplo: blas.Upper,
				Data: make([]float64, l),
			}}
		}
		poolTriDense[i].New = func() interface{} {
			return &TriDense{mat: blas64.Triangular{
				Data: make([]float64, l),
			}}
		}
		poolVecDense[i].New = func() interface{} {
			return &VecDense{mat: blas64.Vector{
				Inc:  1,
				Data: make([]float64, l),
			}}
		}

		poolCDense[i].New = func() interface{} {
			return &CDense{mat: cblas128.General{
				Data: make([]complex128, l),
			}}
		}

		poolFloat64s[i].New = func() interface{} {
			s := make([]float64, l)
			return &s
		}
		poolInts[i].New = func() interface{} {
			s := make([]int, l)
			return &s
		}
	}
}

func getDenseWorkspace(r, c int, clear bool) *Dense { _ = "STUB: not implemented"; return nil }

func putDenseWorkspace(w *Dense) { _ = "STUB: not implemented"; return }

func getSymDenseWorkspace(n int, clear bool) *SymDense { _ = "STUB: not implemented"; return nil }

func putSymDenseWorkspace(s *SymDense) { _ = "STUB: not implemented"; return }

func getTriDenseWorkspace(n int, kind TriKind, clear bool) *TriDense {
	_ = "STUB: not implemented"
	return nil
}

func putTriWorkspace(t *TriDense) { _ = "STUB: not implemented"; return }

func getVecDenseWorkspace(n int, clear bool) *VecDense { _ = "STUB: not implemented"; return nil }

func putVecDenseWorkspace(v *VecDense) { _ = "STUB: not implemented"; return }

func getCDenseWorkspace(r, c int, clear bool) *CDense { _ = "STUB: not implemented"; return nil }

func putCDenseWorkspace(w *CDense) { _ = "STUB: not implemented"; return }

func getFloat64s(l int, clear bool) []float64 { _ = "STUB: not implemented"; return nil }

func putFloat64s(w []float64) { _ = "STUB: not implemented"; return }

func getInts(l int, clear bool) []int { _ = "STUB: not implemented"; return nil }

func putInts(w []int) { _ = "STUB: not implemented"; return }
