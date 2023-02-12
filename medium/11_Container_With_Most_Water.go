/*
You are given an integer array height of length n. There are n vertical
lines drawn such that the two endpoints of the ith line are (i, 0) and
(i, height[i]).

Find two lines that together with the x-axis form a container, such that
the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.



Example 1:


Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array
[1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section)
the container can contain is 49.


Example 2:

Input: height = [1,1]
Output: 1


Constraints:

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	var maxArea float64

	for r > l {
		curArea := float64(r-l) * math.Min(float64(height[l]),
			float64(height[r]))

		maxArea = math.Max(curArea, maxArea)

		if height[l] > height[r] {
			r -= 1
		} else {
			l += 1
		}
	}

	return int(maxArea)
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 4, 3}))

}
