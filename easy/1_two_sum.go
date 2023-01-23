/*
Given an array of integers nums and an integer target, return indices
of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and
you may not use the same element twice.

You can return the answer in any order.


Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]

Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]


Constraints:

2 <= nums.length <= 104
-109 <= nums[i] <= 109
-109 <= target <= 109
Only one valid answer exists.
*/

package main

import "fmt"

/* EXPLANATION
GROUND CASE
- when the array has only 2 values, the sum of the 2 values must be equal
to the target, hence we return the index of the 2 values [0,1]

OTHER CASE
- We create a map called dict. The key of dict will be the difference
between a given value in the array and the target, which the value will
be the index of the value.
- We loop through the loop and search if the num in the array(nums) is a
 key in the map (dict), if yes we return the value of the key and the
current index of array.
- Else we get the difference between the target and the number, we add the
difference to the map as a key with value of the index position.
*/

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	dict := make(map[int]int)

	for i, v := range nums {
		if _, ok := dict[v]; ok {
			return []int{dict[v], i}
		} else {
			diff := target - v
			dict[diff] = i
		}
	}

	return []int{0, 0}
	// The last line is returned to please the go compiler, it does nothing
}

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}
