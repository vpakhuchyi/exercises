package exercises

import (
	"fmt"
	"testing"
)

func Test_SumEvenFibonacci(t *testing.T) {
	fmt.Println(SumEvenFibonacci(1))
}

func SumEvenFibonacci(limit int) (sum int) {
	a, b := 1, 1
	for i := 1; a <= limit; i++ {
		a, b = b, b+a
		if a%2 == 0 {
			sum += a
		}
	}

	return sum
}
