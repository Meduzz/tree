package tree

import (
	"github.com/Meduzz/tree/matchers"
	"github.com/Meduzz/tree/parsers"
)

type (
	// Tree the ... tree
	Tree interface {
		// Add adds a reference to the tree at key
		Add(key, ref string)
		// Remove removes everything at key
		Remove(key string)
		// Lookup tries to find key and return ref or empty string
		Lookup(key string) string
	}

	node struct {
		name     string
		ref      string
		children []*node
	}
)

var (
	// I guess this makes us not thread safe
	p parsers.KeyParser
	m matchers.Matcher
)

// NewTree creates a new tree and allows for some extendability
func NewTree(parser parsers.KeyParser, matcher matchers.Matcher) Tree {
	p = parser
	m = matcher

	return &node{"", "", make([]*node, 0)}
}

func (n *node) Add(key, ref string) {
	keys := p(key)
	old := n
	match := n

	// iterate keys, keep track of the last 2 matching nodes
	for _, k := range keys {
		old = match
		match = match.find(k)

		if match == nil {
			match = &node{k, "", make([]*node, 0)}
			old.children = append(old.children, match)
		}
	}

	match.ref = ref
}

func (n *node) Remove(key string) {
	keys := p(key)
	old := n
	match := n

	// iterate keys, keep track of the last 2 matching nodes
	for _, k := range keys {
		old = match
		match = match.find(k)

		if match == nil {
			return
		}
	}

	keepers := make([]*node, 0)

	for _, m := range old.children {
		if m.name != match.name {
			keepers = append(keepers, m)
		}
	}

	old.children = keepers
}

func (n *node) Lookup(key string) string {
	keys := p(key)
	match := n

	// iterate keys, keep track of the last 2 matching nodes
	for _, k := range keys {
		match = match.find(k)

		if match == nil {
			return ""
		}
	}

	if match == nil {
		return ""
	}

	return match.ref
}

func (n *node) find(key string) *node {
	for _, c := range n.children {
		if m(c.name, key) {
			return c
		}
	}

	return nil
}
