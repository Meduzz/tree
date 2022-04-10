package tree

import (
	"testing"

	"github.com/Meduzz/tree/matchers"
	"github.com/Meduzz/tree/parsers"
)

func TestAdd(t *testing.T) {
	candidate := NewTree(parsers.SplitterParser("/"), matchers.RawMatcher())

	candidate.Add("a/b/c", "c")
	candidate.Add("a/d/e", "e")
	root := candidate.(*node)

	if len(root.children) != 1 {
		t.Errorf("expected # of children of tree to be 1 but was %d", len(root.children))
	}

	a := root.children[0]
	verifyNode(a, "a", "", 2, t)

	b := a.children[0]
	verifyNode(b, "b", "", 1, t)

	d := a.children[1]
	verifyNode(d, "d", "", 1, t)

	c := b.children[0]
	verifyNode(c, "c", "c", 0, t)

	e := d.children[0]
	verifyNode(e, "e", "e", 0, t)
}

func TestRemove(t *testing.T) {
	candidate := NewTree(parsers.SplitterParser("/"), matchers.RawMatcher())

	candidate.Add("a/b/c", "c")
	candidate.Add("a/b/d", "d")
	candidate.Remove("a/b/d")

	root := candidate.(*node)

	if len(root.children) != 1 {
		t.Errorf("expected root.children to be 1 but was %d\n", len(root.children))
	}

	a := root.children[0]
	verifyNode(a, "a", "", 1, t)

	b := a.children[0]
	verifyNode(b, "b", "", 1, t)

	c := b.children[0]
	verifyNode(c, "c", "c", 0, t)

	candidate.Remove("a/b") // remove b and everything above it

	if len(root.children) != 1 {
		t.Errorf("expected root.children to be 1 but was %d\n", len(root.children))
	}

	verifyNode(a, "a", "", 0, t)
}

func TestRemoveNonExistingKey(t *testing.T) {
	candidate := NewTree(parsers.SplitterParser("/"), matchers.RawMatcher())

	candidate.Add("a/b/c", "c")
	candidate.Remove("a/z")

	root := candidate.(*node)

	if len(root.children) != 1 {
		t.Errorf("expected root.children to be 1 but was %d\n", len(root.children))
	}

	a := root.children[0]
	verifyNode(a, "a", "", 1, t)

	b := a.children[0]
	verifyNode(b, "b", "", 1, t)

	c := b.children[0]
	verifyNode(c, "c", "c", 0, t)
}

func TestLookupSomethingReal(t *testing.T) {
	subject := NewTree(parsers.SplitterParser("/"), matchers.RawMatcher())

	subject.Add("a/b/c", "c")
	subject.Add("a/b/d", "d")

	c := subject.Lookup("a/b/c")

	if c != "c" {
		t.Errorf("expected c to be c but was %s\n", c)
	}

	d := subject.Lookup("a/b/d")

	if d != "d" {
		t.Errorf("expected d to be d but was %s\n", d)
	}

	k := subject.Lookup("z/y/x")

	if k != "" {
		t.Errorf("expected k to be '' but was %s\n", k)
	}
}

func verifyNode(node *node, name, ref string, childCount int, t *testing.T) {
	if node.name != name {
		t.Errorf("expected %s.name to be '%s' but was '%s'\n", name, name, node.name)
	}

	if node.ref != ref {
		t.Errorf("expected %s.ref to be '%s' but was '%s'\n", name, ref, node.ref)
	}

	if len(node.children) != childCount {
		t.Errorf("expected %s.children to be %d but was %d\n", name, childCount, len(node.children))
	}
}
