//go:build !debug
// +build !debug

package rdf

type debugger bool

const debug debugger = false

func (d debugger) log(depth int, args ...interface{}) { _ = "STUB: not implemented"; return }
func (d debugger) logf(depth int, format string, args ...interface{}) {
	_ = "STUB: not implemented"
	return
}
func (d debugger) logHashes(depth int, hashes map[string][]byte, size int) {
	_ = "STUB: not implemented"
	return
}
func (d debugger) logParts(depth int, parts byLengthHash) { _ = "STUB: not implemented"; return }
