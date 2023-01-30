/* PLUS ONE
You are given a large integer represented as an integer array digits, where
each digits[i] is the ith digit of the integer. The digits are ordered from
most significant to least significant in left-to-right order. The large integer
does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.



Example 1:

Input: digits = [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Incrementing by one gives 123 + 1 = 124.
Thus, the result should be [1,2,4].
Example 2:

Input: digits = [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
Incrementing by one gives 4321 + 1 = 4322.
Thus, the result should be [4,3,2,2].
Example 3:

Input: digits = [9]
Output: [1,0]
Explanation: The array represents the integer 9.
Incrementing by one gives 9 + 1 = 10.
Thus, the result should be [1,0].


Constraints:

1 <= digits.length <= 100
0 <= digits[i] <= 9
digits does not contain any leading 0's.

*/

package main

import "fmt"

/* EXPLANATION
IDEA - First glance at the problem you'd come to realize the only
issue is dealing with digits that end with 9. Other cases are easy
as you only need to increment the last value by 1.
- Run a while (for) loop over the last number in the array when it
is equal to 9 (the special case).
- In the loop we will make the number at last equal to 0 and update
the value of last by decrementing it i.e last -= 1.
- In the loop we will check for the case when last < 0. This means
all the number in the array/slice were 9s and we need to add 1 to the
front of the array.
- Outside the loop we increment the value of number at last index
i.e digits[last] += 1
*/

func plusOne(digits []int) []int {
	arrOne := []int{1}
	length := len(digits) - 1

	for digits[length] == 9 {
		digits[length] = 0
		length--

		if length < 0 {
			arrOne = append(arrOne, digits...)
			return arrOne
		}

	}

	digits[length] += 1

	return digits

}

func main() {
	fmt.Println(plusOne([]int{9}))
}
