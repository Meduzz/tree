package matchers

// RawMatcher matches the keys against eachother as 2 strings.
func RawMatcher() func(string, string) bool {
	return func(k1, k2 string) bool {
		return k1 == k2
	}
}
