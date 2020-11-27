package exercises

import (
	"fmt"
	"testing"
	"unicode"
)

func Test_Capitalize(t *testing.T) {
	fmt.Println(Capitalize("abcdef", []int{-4,1,2,5,100}))
}


func Capitalize(st string, arr []int) string {
	runes := []rune(st)
	for _, v := range arr {
		if v < 0 || v >= len(runes){
			continue
		}

		runes[v] = unicode.ToUpper(runes[v])
	}

	return string(runes)
}