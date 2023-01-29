/*
Given a sorted array of distinct integers and a target value, return
the index if the target is found. If not, return the index where it
would be if it were inserted in order.

You must write an algorithm with O(log n) runtime complexity.


Example 1:

Input: nums = [1,3,5,6], target = 5
Output: 2
Example 2:

Input: nums = [1,3,5,6], target = 2
Output: 1
Example 3:

Input: nums = [1,3,5,6], target = 7
Output: 4


Constraints:

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums contains distinct values sorted in ascending order.
-104 <= target <= 104
*/

package main

import "fmt"

/* EXPLANATION
IDEA
- We were told that the array is in ascending order (i.e sorted).
Hence we will use binary search algorithm to solve this problem.
- The idea of binary search is to divide the slice/ array in half
with each iteration i.e we compare the mid value in the slice if it
equals that of the target,we return mid; if it is greater than that
of the target, then we know we only want the values which are lesser
than the mid value hence we can reduce our array to nums[:mid] which
is cutting it by half; if the mid value is less than target we reduce to
nums[mid:]. This gives O(logn) time. You can read on binary search.

CODE
- We define two variables
1. high - An integer that stores the index of the highest value in
the current array
2. low - An integer that stores the index of the lowest value in
the current array
- We run a while (for) loop with condition low <= high. This is because
the way we plan on halving our array is by moving the value
of high and low. Also we want to break the loop once high is lesser
than low because it means that the target is not in the array.
- In the for loop we sum up high and low and divide by 2. This gives
us the value of mid.
- We compare mid value (nums[mid]) to the target. 1 of 3 scenarios are
possible
1. They are equal, so we return mid as we have found the position.
2. The mid value is less than target i.e nums[mid] < target, we make
high = mid -1, because we do not need the values above mid ( we know they
they are higher than target).
3. The mid value is higher than target i.e nums[mid] > target, we make
low = mid + 1, because we do not need the values below mid (we know
they are lesser than target).

- Finally, we return high + 1. We only get to this step if the target is
not in array.
*/

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high + 1
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}
