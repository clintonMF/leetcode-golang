/*
Given an integer array nums, find the subarray with the
largest sum, and return its sum.


Example 1:

Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Example 2:

Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Example 3:

Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.


Constraints:

1 <= nums.length <= 105
-104 <= nums[i] <= 104


Follow up: If you have figured out the O(n) solution, try coding another
solution using the divide and conquer approach, which is more subtle.
*/

package main

import "fmt"

/* EXPLANATION

ground case
first case is that of an array with a single value since we were told
len(nums) >= 1.
hence if n == 1 we will return the sole number in the array => num[0]

other case
to solve the problem we need to variables to store the current sum
and maximum sum respectively.
* both variables will be initialized to the value of the first element
1. curSum - saves the current sum of the nums in the array as we loop through
2. maxSum - gives us the max sum of the arrays as we loop through


first we want to run a for loop through the numbers in the array without
the first element (this is because we initialized the value of curSum and maxSum
to that of the first element). we don't want to repeat going over the first element
for the second time.

As we loop through each number we want to compare the number to the sum of the number
and the previous curSum (i.e num to num + curSum). if the number is greater than
the sum of the curSum and num. this means the curSum is negative and thus does not
contribute to value of the maximum number in the array hence we will make the value
of current sum to be the num (i.e curSum = num).
if the sum is greater that means the value of curSum is positive and thus contributes
to the value of current sum so we make curSum = curSum + num.

For maxSum, if the curSum value is greater than that of maxSum we update the value to
that of the curSum, else the value of curSum is less than that of maxSum and we leave it.
*/

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	var maxSum int = nums[0]
	var curSum int = nums[0]

	for _, v := range nums[1:] {
		if curSum+v > v {
			curSum += v
		} else {
			curSum = v
		}

		if curSum > maxSum {
			maxSum = curSum
		}
	}

	return maxSum
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
