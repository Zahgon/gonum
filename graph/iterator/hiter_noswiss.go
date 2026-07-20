//go:build !safe && !go1.24
// +build !safe,!go1.24

package iterator

import "unsafe"

//lint:ignore U1000 This is a verbatim copy of the runtime type.
type hiter struct {
	key         unsafe.Pointer
	elem        unsafe.Pointer
	t           unsafe.Pointer
	h           unsafe.Pointer
	buckets     unsafe.Pointer
	bptr        unsafe.Pointer
	overflow    *[]unsafe.Pointer
	oldoverflow *[]unsafe.Pointer
	startBucket uintptr
	offset      uint8
	wrapped     bool
	B           uint8
	i           uint8
	bucket      uintptr
	checkBucket uintptr
}

func (h *hiter) initialized() bool { _ = "STUB: not implemented"; return false }
