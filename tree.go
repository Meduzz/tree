package tree

import (
	"sync"

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
		matcher  matchers.Matcher
	}

	root struct {
		*node
		lock   *sync.Mutex
		parser parsers.KeyParser
	}
)

// NewTree creates a new tree and allows for some extendability
func NewTree(parser parsers.KeyParser, matcher matchers.Matcher) Tree {
	return &root{
		node:   &node{"", "", make([]*node, 0), matcher},
		lock:   &sync.Mutex{},
		parser: parser,
	}
}

func (n *root) Add(key, ref string) {
	n.lock.Lock()
	defer n.lock.Unlock()

	keys := n.parser(key)
	old := n.node
	match := n.node

	// iterate keys, keep track of the last 2 matching nodes
	for _, k := range keys {
		old = match
		match = match.find(k)

		if match == nil {
			match = &node{k, "", make([]*node, 0), n.matcher}
			old.children = append(old.children, match)
		}
	}

	match.ref = ref
}

func (n *root) Remove(key string) {
	n.lock.Lock()
	defer n.lock.Unlock()

	keys := n.parser(key)
	old := n.node
	match := n.node

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

func (n *root) Lookup(key string) string {
	keys := n.parser(key)
	match := n.node

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
		if n.matcher(c.name, key) {
			return c
		}
	}

	return nil
}
