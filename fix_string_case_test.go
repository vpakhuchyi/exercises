package exercises

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func Test_solve(t *testing.T) {
	fmt.Println(solve("CODe"))
}

func solve(str string) string {
	runesLen := len([]rune(str))
	if runesLen == 1 { // one letter string doesn't need a transformation
		return str
	}

	var upperRunes int
	for _, rn := range str {
		if unicode.IsUpper(rn) {
			upperRunes++
		}
	}

	if upperRunes > runesLen/2 {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
