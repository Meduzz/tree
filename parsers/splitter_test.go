package parsers

import "testing"

func TestSplittingStuff(t *testing.T) {
	expected := []string{"a", "b", "c"}
	subject := SplitterParser("/")
	res := subject("a/b/c")

	if len(res) != 3 {
		t.Errorf("expected array to be of length 3 but was %d\n", len(res))
	}

	if res[0] != expected[0] && res[1] != expected[1] && res[2] != expected[2] {
		t.Errorf("expected array to be a,b,c but was %v", res)
	}
}
