package exercises

import (
	"fmt"
	"testing"
)

func Test_Gps(t *testing.T) {
	fmt.Println(Gps(20, []float64{0.0, 0.23, 0.46, 0.69, 0.92, 1.15, 1.38, 1.61}))
}

func Gps(s int, x []float64) int {
	var maxSpeed float64
	for i, v := range x {
		// avoid slice bounds exceeding
		if i+1 >= len(x) {
			return int(maxSpeed)
		}

		if speed := (3600 * (x[i+1] - v)) / float64(s); speed > maxSpeed {
			maxSpeed = speed
		}
	}

	return int(maxSpeed)
}