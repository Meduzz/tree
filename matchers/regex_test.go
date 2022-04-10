package matchers

import "testing"

func TestSomeMatching(t *testing.T) {
	subject := RegexMatcher("*", UrlSafe)

	// not a wildcard
	if !subject("k1", "k1") {
		t.Error("expected 2 equal strings to match")
	}

	if subject("k1", "k2") {
		t.Error("expected 2 different strings to not match")
	}

	if subject("k1", "K1") {
		t.Error("expected matcher to be case sensitive")
	}

	// wildcards
	if !subject("*", "k1") {
		t.Error("expected wildcard to match k1")
	}

	if !subject("*", "a_lot-of-sp3ciul_ch4r5") {
		t.Error("expected wildcard to match speciul text")
	}

	if subject("*", "!@+=*") {
		t.Error("did not expect wildcard to match unsafe chars")
	}
}
