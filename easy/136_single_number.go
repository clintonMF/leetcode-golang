/*
Given a non-empty array of integers nums, every element appears twice
except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use
only constant extra space.



Example 1:

Input: nums = [2,2,1]
Output: 1

Example 2:

Input: nums = [4,1,2,1,2]
Output: 4

Example 3:

Input: nums = [1]
Output: 1


Constraints:

1 <= nums.length <= 3 * 104
-3 * 104 <= nums[i] <= 3 * 104
Each element in the array appears twice except for one element which appears
only once.
*/

package main

import (
	"fmt"
	"sort"
)

/* EXPLANATION
IDEA 1
- loop over the array and save the elements in a map: the key is the element
and the value is the number of times it occurs.
- loop over the map and return the key where value is equal to 1
- This solution is easy, intuitive and works fine. but it uses O(N) space
the question asked us to implement it in O(1) space i.e constant space

IDEA 2
- Sort the array
- If the array is sorted and we expect each element except one to be
duplicated. We should also expect that every element in an even number
index position should be the same element in the next odd number position
i.e [1,1,2,] the first 1 is in at index 0, we expect to see 1 again at
index 1. If This is not the case we immediately return the number because
we are sure it is the number that is not repeated.
* NOTE
* You can't run your loop to the end this is to avoid index out of range
* error
* The way around this is to run it to the second to the last element
* If the loop runs without returning the number then we know the single
* number must be the last element and we return it.
-
*/

// 1

// func singleNumber(nums []int) int {
// 	var dict map[int]int = make(map[int]int)

// 	for _, v := range nums {
// 		if _, ok := dict[v]; ok {
// 			delete(dict, v)
// 		} else {
// 			dict[v] = 1
// 		}
// 	}

// 	for i, v := range dict {
// 		if v == 1 {
// 			return i
// 		}
// 	}

// 	return -1

// }

// 2

func singleNumber(nums []int) int {
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i%2 == 0 && nums[i] != nums[i+1] {
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 4, 6, 4}))
}
