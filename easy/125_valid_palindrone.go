/* A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.


Constraints:

1 <= s.length <= 2 * 105
s consists only of printable ASCII characters.
*/

package main

/* EXPLANATION
 */

import (
	"fmt"
	"regexp"
	"strings"
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
- you can check the ground case, which is when the number is 1 character or
noe, you return true (1 character is a palindrome).
- We clean the the string by first removing all invalid characters (i.e space),
then we take it to lower case.
- We create 2 variables
left - equal to the first character index = 0
right - last character index - len(word) - 1

- We run a loop while left <= right, in the loop if their characters are
not equal we return false, else we increment left and decrement right.
- if the loop completes running all left and right characters are equal and
we return true
*/

func isPalindrome(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = regexp.MustCompile(`[^a-zA-Z0-9]+`).ReplaceAllString(s, "")
	lower := strings.ToLower(s)

	left := 0
	right := len(lower) - 1

	for left < right {
		if lower[left] != lower[right] {
			return false
		}
		left = left + 1
		right = right - 1
	}

	return true

}

func main() {
	str := "HER i am to say"
	str2 := "A man, a plan, a canal: Panama"

	fmt.Println(isPalindrome(str))
	fmt.Println(isPalindrome(str2))
}
