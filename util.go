package valgo

import (
	"encoding/json"
	"sort"
	"strings"
	"unicode"
)

func concatString(stringA string, stringB string) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString(stringA)
	strBuilder.WriteString(stringB)
	return strBuilder.String()
}

func humanizeName(name string) string {
	in := []rune(strings.TrimSpace(name))
	space := []rune(" ")[0]
	lastIndex := len(in) - 1

	out := strings.Builder{}

	for i, c := range in {
		if i == 0 {
			if unicode.IsLower(c) {
				out.WriteRune(unicode.ToUpper(c))
			} else {
				out.WriteRune(c)
			}
		} else {
			cb := in[i-1]
			if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
				if !unicode.IsLetter(cb) && !unicode.IsNumber(cb) {
					continue
				} else {
					out.WriteRune(space)
				}
			} else if unicode.IsUpper(c) {
				isLast := i == lastIndex
				var cn rune
				if !isLast {
					cn = in[i+1]
				}
				if unicode.IsUpper(cb) && (isLast || (unicode.IsUpper(cn) || !unicode.IsLetter(cn))) {
					out.WriteRune(c)
				} else {
					if unicode.IsLetter(cb) || unicode.IsNumber(cb) {
						out.WriteRune(space)
					}
					if !unicode.IsUpper(cb) && (!isLast && unicode.IsUpper(cn)) {
						out.WriteRune(c)
					} else {
						out.WriteRune(unicode.ToLower(c))
					}
				}
			} else if unicode.IsNumber(c) {
				if unicode.IsLetter(cb) {
					out.WriteRune(space)
				}
				out.WriteRune(c)
			} else {
				out.WriteRune(unicode.ToLower(c))
			}
		}
	}

	return out.String()
}

// serializes the error messages into sorted JSON for consistency in
// documentation examples.
func sortedErrorMarshalForDocs(e *Error) ([]byte, error) {
	// Create a slice to hold the keys from the map, so we can sort them.
	keys := make([]string, 0, len(e.errors))
	for k := range e.errors {
		keys = append(keys, k)
	}
	sort.Strings(keys) // Sort the keys alphabetically.

	// Create a map to hold the sorted key-value pairs.
	sortedErrors := make(map[string]interface{}, len(keys))
	for _, k := range keys {
		messages := make([]string, len(e.errors[k].Messages()))
		copy(messages, e.errors[k].Messages())
		sort.Strings(messages)
		sortedErrors[k] = messages
	}

	// Marshal the sorted map to JSON.
	return json.Marshal(sortedErrors)
}
