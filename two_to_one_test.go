package exercises

import (
	"fmt"
	"sort"
	"testing"
)

func Test_TwoToOne(t *testing.T) {
	fmt.Println(TwoToOne("xyaabbbccccdefww", "xxxxyyyyabklmopq"))
}

func TwoToOne(s1 string, s2 string) string {
	set := map[rune]struct{}{}

	for _, v := range s1 + s2 {
		set[v] = struct{}{}
	}

	buf := make([]rune, 0, len(set))
	for rn := range set {
		buf = append(buf, rn)
	}

	sort.Slice(buf, func(i, j int) bool {
		return buf[i] < buf[j]
	})

	return string(buf)
}
