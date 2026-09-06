//go:build !bounds
// +build !bounds

package mat

func (m *Dense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (m *Dense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (m *Dense) Set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (m *Dense) set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (m *CDense) At(i, j int) complex128 { _ = "STUB: not implemented"; return 0 }

func (m *CDense) at(i, j int) complex128 { _ = "STUB: not implemented"; return 0 }

func (m *CDense) Set(i, j int, v complex128) { _ = "STUB: not implemented"; return }

func (m *CDense) set(i, j int, v complex128) { _ = "STUB: not implemented"; return }

func (v *VecDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) AtVec(i int) float64 { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) at(i int) float64 { _ = "STUB: not implemented"; return 0 }

func (v *VecDense) SetVec(i int, val float64) { _ = "STUB: not implemented"; return }

func (v *VecDense) setVec(i int, val float64) { _ = "STUB: not implemented"; return }

func (s *SymDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymDense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymDense) SetSym(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (s *SymDense) set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (t *TriDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriDense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriDense) SetTri(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (t *TriDense) set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (b *BandDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (b *BandDense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (b *BandDense) SetBand(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (b *BandDense) set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (s *SymBandDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymBandDense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (s *SymBandDense) SetSymBand(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (s *SymBandDense) set(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (t *TriBandDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriBandDense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (t *TriBandDense) SetTriBand(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (d *DiagDense) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (d *DiagDense) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (d *DiagDense) SetDiag(i int, v float64) { _ = "STUB: not implemented"; return }

func (d *DiagDense) setDiag(i int, v float64) { _ = "STUB: not implemented"; return }

func (a *Tridiag) At(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (a *Tridiag) at(i, j int) float64 { _ = "STUB: not implemented"; return 0 }

func (a *Tridiag) SetBand(i, j int, v float64) { _ = "STUB: not implemented"; return }

func (a *Tridiag) set(i, j int, v float64) { _ = "STUB: not implemented"; return }
