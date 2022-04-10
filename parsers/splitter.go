package parsers

import "strings"

func SplitterParser(char string) KeyParser {
	return func(key string) []string {
		return strings.Split(key, char)
	}
}
