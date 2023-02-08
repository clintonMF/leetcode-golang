/*
Given an integer array nums, return an array answer such that answer[i]
is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a
32-bit integer.

You must write an algorithm that runs in O(n) time and without using
the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]
Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]


Constraints:

2 <= nums.length <= 105
-30 <= nums[i] <= 30
The product of any prefix or suffix of nums is guaranteed to fit in a
32-bit integer.


Follow up: Can you solve the problem in O(1) extra space complexity?
(The output array does not count as extra space for space complexity
analysis.)
*/

package main

import "fmt"

func productExceptSelf(nums []int) []int {

	if len(nums) == 2 {
		return []int{nums[1], nums[0]}
	}

	res := []int{}
	pre := 1
	res = append(res, pre)

	for i := 0; i < len(nums)-1; i++ {
		res = append(res, 1)
		pre *= nums[i]
		res[i+1] *= pre
	}
	fmt.Println(res)

	post := 1

	for i := len(nums) - 1; i > 0; i-- {
		post *= nums[i]
		res[i-1] *= post
	}

	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
