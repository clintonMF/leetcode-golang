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

/* EXPLANATION
IDEA
- The idea utilizes the sliding window concept
- we will create a map/hashset to store the characters of the string
- we can loop through the characters using the index.
	if we find the given character in the map. We remove characters of
	the string from the dictionary until the character isn't in the map.
	this is done by keeping track of the first member of the dict by
	an integer l.
	else, we add the character to the map and check if the current length
	is the maximum.

CODE
- compute the function for finding max of 2 integers
- Initialize the map called charSet, it uses byte as key and
bool as value

*NOTE: byte as key is done because golang returns byte when you
* get a character by its index from a string. Since each character
* has a unique byte there is no need to change it to string.
*NOTE: It is imperative to use bool data type for the value: it helps
* with the condition later on.

- Initialize l: an integer that tracks the current window (pointer)
position
- Initialize maxLength: an integer that tracks the maximum length

- Run a for loop on the string keeping/using the index
	- run a while loop (For, golang doesn't have a while loop) with
	condition if the current character is present in the map.
	- If true, remove the element s[l] from the dict and increment l.
	- else, add the character to the map and check if the current
	length is the max length.

- return maxLength
*/

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
