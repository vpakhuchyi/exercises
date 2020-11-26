package exercises

import (
	"fmt"
	"testing"
)

func Test_CartesianNeighbor(t *testing.T) {
	fmt.Println(CartesianNeighbor(2, 2))
}

/*
	Possible dots:

	[x-1;y+1][x;y+1][x+1;y+1]
	[ x-1;y ][  *  ][ x+1;y ]
	[x-1;y-1][x;y-1][x+1;y-1]
*/
func CartesianNeighbor(x, y int) [][]int {
	return [][]int{
		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
		{x - 1, y},
		{x + 1, y},
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},
	}
}
