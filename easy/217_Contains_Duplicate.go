/*
Given an integer array nums, return true if any value appears at least
twice in the array, and return false if every element is distinct.

Example 1:

Input: nums = [1,2,3,1]
Output: true

Example 2:

Input: nums = [1,2,3,4]
Output: false

Example 3:

Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true


Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109

*/

package main

import "fmt"

/* EXPLANATION
GROUND CASE
- if length of the array is less than 2, then no value in the
array can be repeated, hence we return false.

OTHER CASE
- We need to create a variable dict (i.e it is a map) which will store the values in the
array as keys.
- we loop through the values in the arrays. We check if the value is in
the map (dict). if yes we immediately return true because we have seen a repeated value.
Else we add the value to the map.
- if after running the loop we get no repeated value, we return false.
*/

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}

	dict := make(map[int]int)

	for _, v := range nums {
		if _, ok := dict[v]; ok {
			return true
		} else {
			dict[v] = 1
		}
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{2, 3, 7, 8, 2}))
}
