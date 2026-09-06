//go:build !safe && go1.24
// +build !safe,go1.24

package iterator

import "unsafe"

//lint:ignore U1000 This is a verbatim copy of the runtime type.
type hiter struct {
	key         unsafe.Pointer
	elem        unsafe.Pointer
	typ         unsafe.Pointer
	m           unsafe.Pointer
	entryOffset uint64
	dirOffset   uint64
	clearSeq    uint64
	globalDepth uint8
	dirIdx      int
	tab         unsafe.Pointer
	group       groupReference
	entryIdx    uint64
}

//lint:ignore U1000 This is a verbatim copy of the runtime type.
type groupReference struct {
	data unsafe.Pointer
}

func (h *hiter) initialized() bool { _ = "STUB: not implemented"; return false }
