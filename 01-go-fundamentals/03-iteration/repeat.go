package iteration

import "strings"

func Repeat(character string, iterations int) string {
	var repeated strings.Builder
	for i := 0; i < iterations; i++ {
		repeated.WriteString(character)
	}
	return repeated.String()
}
