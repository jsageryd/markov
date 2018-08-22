package blacklist

import (
	"crypto/sha1"
	"encoding/gob"
	"hash"
	"io"
)

// Blacklist implements a blacklist wuth support for adding items and checking
// whether they exist.
type Blacklist struct {
	items map[string]struct{}
	hash  hash.Hash
}

// New returns a new Blacklist
func New() *Blacklist {
	return &Blacklist{
		items: make(map[string]struct{}),
		hash:  sha1.New(),
	}
}

// Add adds the given item to the blacklist.
func (b *Blacklist) Add(item []byte) {
	b.items[string(b.hashItem(item))] = struct{}{}
}

// Has returns true if the given item is in the list, otherwise false.
func (b *Blacklist) Has(item []byte) bool {
	_, ok := b.items[string(b.hashItem(item))]
	return ok
}

func (b *Blacklist) hashItem(item []byte) []byte {
	b.hash.Write(item)
	h := b.hash.Sum(nil)
	b.hash.Reset()
	return h[:len(h)/2]
}

// ExportState exports the state of the list.
func (b *Blacklist) ExportState(w io.Writer) error {
	return gob.NewEncoder(w).Encode(b.items)
}

// ImportState imports the given state.
func (b *Blacklist) ImportState(r io.Reader) error {
	return gob.NewDecoder(r).Decode(&b.items)
}
