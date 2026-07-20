package rdf

import (
	"hash"
)

func Isomorphic(a, b []*Statement, decomp bool, h hash.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

func lexicalHashes(dst [][]byte, hashes map[string][]byte) { _ = "STUB: not implemented"; return }

func IsoCanonicalHashes(statements []*Statement, decomp, dist bool, h hash.Hash, zero []byte) (hashes map[string][]byte, terms map[string]map[string]bool) {
	_ = "STUB: not implemented"
	return nil, nil
}

func C14n(dst, src []*Statement, terms map[string]map[string]bool) ([]*Statement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func translate(term string, mapping map[string]string) string { _ = "STUB: not implemented"; return "" }

func sortC14nStatements(statements []*Statement) { _ = "STUB: not implemented"; return }

func hashBNodes(statements []*Statement, h hash.Hash, zero []byte, hash0 map[string][]byte) (hash *table, disjoint bool) {
	_ = "STUB: not implemented"
	return nil, false
}

type table struct {
	hashOf map[string][]byte

	termsFor map[string]map[string]bool

	isBlank map[string]bool

	blanks []string
}

func newTable() *table { _ = "STUB: not implemented"; return nil }

func (t *table) wasCloned() bool { _ = "STUB: not implemented"; return false }

func (t *table) isNew() bool { _ = "STUB: not implemented"; return false }

func (t *table) clone() *table { _ = "STUB: not implemented"; return nil }

func (t *table) set(term string, hash []byte) { _ = "STUB: not implemented"; return }

func (t *table) allUnique() bool { _ = "STUB: not implemented"; return false }

func (t *table) changedFrom(last *table) bool { _ = "STUB: not implemented"; return false }

type hashBag struct {
	hash      hash.Hash
	hashesFor map[string][][]byte
}

func newHashBag(h hash.Hash, t *table) hashBag { _ = "STUB: not implemented"; return *new(hashBag) }

func (b hashBag) add(term string, hash []byte) { _ = "STUB: not implemented"; return }

func (b hashBag) sum(term string) []byte { _ = "STUB: not implemented"; return nil }

func hashTuple(h hash.Hash, t ...[]byte) []byte { _ = "STUB: not implemented"; return nil }

func hashBNodesPerSplit(statements []*Statement, decomp bool, h hash.Hash, zero []byte) (hash *table, parts byLengthHash, disjoint bool) {
	_ = "STUB: not implemented"
	return nil, *new(byLengthHash), false
}

func appendOrdered(parts byLengthHash, partSets map[string]map[string]bool) byLengthHash {
	_ = "STUB: not implemented"
	return *new(byLengthHash)
}

type byLengthHash struct {
	nodes [][]string

	hashes []string
}

func (s byLengthHash) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s byLengthHash) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s byLengthHash) Swap(i, j int) { _ = "STUB: not implemented"; return }

func allUnique(hashes map[string][]byte) bool { _ = "STUB: not implemented"; return false }

func split(statements []*Statement) [][]*Statement { _ = "STUB: not implemented"; return nil }

func distinguish(statements []*Statement, dist bool, h hash.Hash, zero []byte, hash *table, parts byLengthHash, lowest map[string][]byte, depth int) map[string][]byte {
	_ = "STUB: not implemented"
	return nil
}

func graphLess(statements []*Statement, a, b map[string][]byte) bool {
	_ = "STUB: not implemented"
	return false
}

type lexicalStatements struct {
	statements []*Statement
	hashes     map[string][]byte
}

func newLexicalStatements(statements []*Statement, hash map[string][]byte) lexicalStatements {
	_ = "STUB: not implemented"
	return *new(lexicalStatements)
}

func sub(a, b lexicalStatements, max int) []*Statement { _ = "STUB: not implemented"; return nil }

func (s lexicalStatements) Len() int           { _ = "STUB: not implemented"; return 0 }
func (s lexicalStatements) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s lexicalStatements) Swap(i, j int) { _ = "STUB: not implemented"; return }

type relabeledStatement struct {
	statement *Statement
	labels    map[string][]byte
}

func (a relabeledStatement) less(b relabeledStatement) bool {
	_ = "STUB: not implemented"
	return false
}

func (s relabeledStatement) String() string { _ = "STUB: not implemented"; return "" }

type relabeledTerm struct {
	term   Term
	labels map[string][]byte
}

func (a relabeledTerm) less(b relabeledTerm) bool { _ = "STUB: not implemented"; return false }

func unquoteIRI(s string) string { _ = "STUB: not implemented"; return "" }

func (t relabeledTerm) String() string { _ = "STUB: not implemented"; return "" }
