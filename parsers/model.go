package parsers

type (
	// KeyParser a function that turns a single string into many if needed.
	KeyParser func(string) []string
)
