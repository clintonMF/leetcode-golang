/*
Given a string s, find the length of the sub_arr substring
without repeating characters.

Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

Constraints:

0 <= s.length <= 5 * 104
s consists of English letters, digits, symbols and spaces.
*/
package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func lengthOfLongestSubstring(s string) int {
	charSet := make(map[byte]bool)
	l := 0
	maxLength := 0

	for i := 0; i < len(s); i++ {
		for charSet[s[i]] {
			delete(charSet, s[l])
			l += 1
		}
		charSet[s[i]] = true
		maxLength = max(maxLength, i-l+1)
	}

	return maxLength
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbcbb"))
}
