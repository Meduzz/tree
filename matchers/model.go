package matchers

type (
	// Matcher a function to controll if 2 strings matches
	Matcher func(string, string) bool
)
