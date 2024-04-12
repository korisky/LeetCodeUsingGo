package Ex42_TrappingRainWater

import (
	"fmt"
	"testing"
)

func trap(height []int) int {

	var l2r = make([]int, len(height))
	var r2l = make([]int, len(height))

	l2r[0] = height[0]
	r2l[len(height)-1] = height[len(height)-1]

	for i := 1; i < len(height); i++ {
		l2r[i] = max(l2r[i-1], height[i])
		j := len(height) - 1 - i
		r2l[j] = max(r2l[j+1], height[j])
	}

	water := 0
	for i := 0; i < len(height); i++ {
		water += int(min(l2r[i], r2l[i])) - height[i]
	}
	return water
}

func Test_trap(t *testing.T) {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
