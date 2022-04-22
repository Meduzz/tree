package matchers

import (
	"regexp"
)

// safe..ish chars for urls
const UrlSafe = "[a-zA-Z0-9-_%.!]+"

// RegexMatcher first checks if k1 == wildcard, then uses regex to match on k2, otherwise checks k1 == k2.
func RegexMatcher(wildcard, regex string) func(string, string) bool {
	r, err := regexp.Compile(regex)

	if err != nil {
		// im sorry
		panic(err)
	}

	return func(k1, k2 string) bool {
		if k1 == wildcard {
			return r.MatchString(k2)
		} else {
			return k1 == k2
		}
	}
}
