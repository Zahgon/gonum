package rdf

import (
	"hash"

	"gonum.org/v1/gonum/stat/combin"
)

func Deduplicate(s []*Statement) []*Statement { _ = "STUB: not implemented"; return nil }

func isSameStatement(a, b *Statement) bool { _ = "STUB: not implemented"; return false }

func URGNA2012(dst, src []*Statement) ([]*Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func URDNA2015(dst, src []*Statement) ([]*Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type urna struct {
	canon *issuer

	hashes map[string]string

	statementsFor map[string][]*Statement

	hash hash.Hash

	hashToRelated relatedHashCreator

	label string
}

func (u *urna) relabel(dst, src []*Statement) ([]*Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lexicallySortedPathHashes(paths map[string][]*issuer) []string {
	_ = "STUB: not implemented"
	return nil
}

func translateURNA(term string, mapping map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func (u *urna) hashFirstDegreeQuads(b string) string { _ = "STUB: not implemented"; return "" }

func replaceBlank(b, matching, label string) string { _ = "STUB: not implemented"; return "" }

func (u *urna) hashNDegreeQuads(b string, names *issuer) ([]byte, *issuer) {
	_ = "STUB: not implemented"
	return nil, nil
}

func lexicallySortedTermHashes(termsFor map[string][]string) []string {
	_ = "STUB: not implemented"
	return nil
}

type relatedHashCreator func(b string, names *issuer) map[string][]string

func (u *urna) hashToRelatedURDNA2015(b string, names *issuer) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (u *urna) hashToRelatedURGNA2012(b string, names *issuer) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func (u *urna) hashRelatedBlank(term string, s *Statement, names *issuer, pos byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

type issuer struct {
	prefix  string
	issued  map[string]string
	ordered []string
}

func newIssuer(prefix string) *issuer { _ = "STUB: not implemented"; return nil }

func (i *issuer) issueFor(b string) string { _ = "STUB: not implemented"; return "" }

func (i *issuer) has(id string) bool { _ = "STUB: not implemented"; return false }

func (i *issuer) clone() *issuer { _ = "STUB: not implemented"; return nil }

func hex(data []byte) []byte { _ = "STUB: not implemented"; return nil }

type permutations struct {
	src  []string
	dst  []string
	idx  []int
	perm *combin.PermutationGenerator
}

func newPermutations(src []string) *permutations { _ = "STUB: not implemented"; return nil }

func (p *permutations) next() bool { _ = "STUB: not implemented"; return false }

func (p *permutations) permutation() []string { _ = "STUB: not implemented"; return nil }
