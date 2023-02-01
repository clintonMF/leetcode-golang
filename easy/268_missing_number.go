/*
Given an array nums containing n distinct numbers in the range [0, n],
return the only number in the range that is missing from the array.



Example 1:

Input: nums = [3,0,1]
Output: 2
Explanation: n = 3 since there are 3 numbers, so all numbers are in the range [0,3].
2 is the missing number in the range since it does not appear in nums.
Example 2:

Input: nums = [0,1]
Output: 2
Explanation: n = 2 since there are 2 numbers, so all numbers are in the range [0,2].
2 is the missing number in the range since it does not appear in nums.
Example 3:

Input: nums = [9,6,4,2,3,5,7,0,1]
Output: 8
Explanation: n = 9 since there are 9 numbers, so all numbers are in the range [0,9].
8 is the missing number in the range since it does not appear in nums.


Constraints:

n == nums.length
1 <= n <= 104
0 <= nums[i] <= n
All the numbers of nums are unique.


Follow up: Could you implement a solution using only O(1) extra space complexity
and O(n) runtime complexity?
*/

package main

import "fmt"

/*
The logic behind the solution below lies in that fact that if you sum
all the numbers from 1 to n and subtract it from the sum of numbers in
the list/slice, this will give us the missing number in the list/slice.

knowing this we can run a for loop over the slice to get the total sum of
numbers in the list. we can also run a for loop with n being the highest
number to get the total sum of numbers from 1 to n. then take the difference

but this will require 2 for loops, luckily for us we can combine both loops
and take the difference on each iteration.
To do this we will need 3 variables
1. the length of the list which is also n
2. Tracker - This is the variable that computes the sum of the values from
n to 1. e.g if we were to create a for loop to get the value sum from 1 to n
it will have us add the value of n as we decrease n on each iteration.
i.e
  sum = sum + tracker
  tracker = tracker - 1

3. diff - This variable keeps the value of the difference between the sum of 1 to n and sum of list
for each iteration we add the previous diff to the tracker and the current num in the slice.
i.e diff = diff + target - nums[i]


example: given nums [2,0,1]
n	diff	tracker		nums[i]
0	 1	   	3			2
1	 3		2			0
2	 1		1 			3

we return the final diff which is 1. the missing digit in the example is 1

*/

func missingNumber(nums []int) int {
	n := len(nums)
	tracker := n
	diff := 0

	for i := 0; i < n; i++ {
		diff = diff + tracker - nums[i]
		tracker -= 1
	}

	return diff

}

func main() {
	fmt.Println(missingNumber([]int{4, 2, 3, 0, 1, 6}))
}
