/*
You are given a string s and an integer k. You can choose any
character of the string and change it to any other uppercase
English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same
letter you can get after performing the above operations.

Example 1:

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Example 2:

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.

Constraints:

1 <= s.length <= 105
s consists of only uppercase English letters.
0 <= k <= s.length
*/
package main

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func characterReplacement(s string, k int) int {
	count := make(map[byte]int)
	maxF := 0
	l := 0
	res := 0

	for r := 0; r < len(s); r++ {
		count[s[r]] += 1
		maxF = max(maxF, count[s[r]])

		if r-l+1-maxF > k {
			count[s[l]] -= 1
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}
