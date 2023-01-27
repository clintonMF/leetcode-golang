/*
Given a 1-indexed array of integers numbers that is already sorted
in non-decreasing order, find two numbers such that they add up to
a specific target number. Let these two numbers be numbers[index1]
and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

Return the indices of the two numbers, index1 and index2, added by
one as an integer array [index1, index2] of length 2.

The tests are generated such that there is exactly one solution.
You may not use the same element twice.

Your solution must use only constant extra space.

Example 1:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1,
index2 = 2. We return [1, 2].

Example 2:

Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1,
index2 = 3. We return [1, 3].

Example 3:

Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1,
index2 = 2. We return [1, 2].

Constraints:

2 <= numbers.length <= 3 * 104
-1000 <= numbers[i] <= 1000
numbers is sorted in non-decreasing order.
-1000 <= target <= 1000
The tests are generated such that there is exactly one solution.
*/
package main

import "fmt"

/* EXPLANATION
It is very possible to solve this problem with the solution of problem
"1_two_sum". You will only need to add an extra 1 to the returned values
because this problems index starts from 1 not 0.
But we will  solve this problem in a different way because we are
guaranteed that the array is sorted. Also this new solution will be helpful
in understanding the solution to another question (3sum closest).

IDEA
- Given that the array is sorted, we are aware that the numbers at the end
(right) are higher than those at the start (left).
- When we add the first and last numbers if they are greater than the target
that must mean that the last number is too high and hence we move to the
second to the last number by decrementing its index. if they are lesser than
the target that must mean that the first number is too small and hence we
move to the second number.
- repeating the process above will lead us to the numbers that sum up to the
target.
- we can return their index and add 1 to it.
*/

func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		if numbers[right]+numbers[left] < target {
			left += 1
		} else if numbers[right]+numbers[left] > target {
			right -= 1
		} else {
			return []int{left + 1, right + 1}
		}
	}

	return []int{0, 0}
	// the code above will never run as we are guaranteed to get a result,
	// it is only there to satisfy the compiler

}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 13}, 9))
}
