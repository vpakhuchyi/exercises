package exercises

import (
	"fmt"
	"testing"
)

func Test_WordsToMarks(t *testing.T) {
	fmt.Println(WordsToMarks("attitude"))
}

func WordsToMarks(s string) (result int) {
	mapping := createMapping()
	for _, rn := range s {
		result += mapping[rn]
	}

	return result
}

const aIndex = 'a' // index of first eng lowercase character
const engLettersNumber = 26

func createMapping() map[rune]int {
	mapping := make(map[rune]int, 26)
	for c, i := 1, aIndex; i < aIndex+engLettersNumber; i++ {
		mapping[i] = c
		c++
	}

	return mapping
}

//better solution made by khsadira
//func WordsToMarks(s string) int {
//	count := 0
//	for _, i := range s {
//		count += int(i) - 'a' + 1
//	}
//	return count
//}
