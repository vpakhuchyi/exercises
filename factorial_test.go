package exercises

import (
	"fmt"
	"testing"
)

func Test_Factorial(t *testing.T) {
	fmt.Println(Factorial(4))
}

func Factorial(n int) int {
	result := 1
	for ; n > 0; n-- {
		result *= n
	}

	return result
}
