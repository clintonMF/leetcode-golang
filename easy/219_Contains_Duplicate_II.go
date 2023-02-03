/*
Given an integer array nums and an integer k, return true if there are
two distinct indices i and j in the array such that nums[i] == nums[j]
and abs(i - j) <= k.

Example 1:

Input: nums = [1,2,3,1], k = 3
Output: true
Example 2:

Input: nums = [1,0,1,1], k = 1
Output: true
Example 3:

Input: nums = [1,2,3,1,2,3], k = 2
Output: false

Constraints:

1 <= nums.length <= 105
-109 <= nums[i] <= 109
0 <= k <= 105
*/
package main

/* EXPLANATION
GROUND CASE
- If you have just one integers return false (you can't have 2 integers
in an array containing one integer).

MAIN IDEA
- We will loop through the array/slice and store the integers and index as
key and value respectively in an array.
- When we get to an integer thats in the dict we confirm if the difference
between the current index and the index of the other is less than or equal
to k. If it is we return true, else we update the current index.
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	dict := make(map[int]int)

	for i, v := range nums {
		if _, ok := dict[v]; ok {
			diff := i - dict[v]

			if k >= diff {
				return true
			} else {
				dict[v] = i
			}
		} else {
			dict[v] = i
		}
	}

	return false
}
