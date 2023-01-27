/*
Given an integer x, return true if x is a
palindrome, and false otherwise.



Example 1:

Input: x = 121
Output: true
Explanation: 121 reads as 121 from left to right and from right to left.
Example 2:

Input: x = -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: x = 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.


Constraints:

-231 <= x <= 231 - 1

*/

package main

import (
	"fmt"
	"strconv"
)

/* EXPLANATION

IDEA
- A palindrome is a word that reads the same when spelt backwards.
this also means that the first character is the same as the last character,
second character and second to the last are also the same and so on.
- Hence, we can create two pointers, the first will start at the first
character and we will increment it. The second will start at the last
character and we will decrement.
- If the word is a palindrome, both pointers will have the same character,
- if both are not the same the word is not a palindrome.

CODE
- we start by converting the number to string; this will enable us perform
some needed operations.
- you can check the ground case, which is when the number is 1 character or
noe, you return true (1 character is a palindrome).
- We create 2 variables
left - equal to the first character index = 0
right - last character index - len(word) - 1

- we run a loop while left <= right, in the loop if their characters are
not equal we return false, else we increment left and decrement right.
- if the loop completes running all left and right characters are equal and
we return true
*/

func isPalindrom(x int) bool {
	str := strconv.Itoa(x)

	if len(str) < 2 {
		return true
	}

	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	fmt.Println(isPalindrom(121))
	fmt.Println(isPalindrom(-121))
	fmt.Println(isPalindrom(10))

}
