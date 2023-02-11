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

/* EXPLANATION
IDEA
- Using the sliding window technique we can add the characters of
the string into a map; keeping count of the number of times a certain
value occurs.
- We want the longest string with the same character and we are told
that we can have k number of different characters in between.
i.e if k = 1 then "abaa" is a valid because b appears just once.
- With this in mind we realize that the only value in the map we are
interested in is that of the highest occurrence, using the example given
above we will only be interested in the number of times "a" occurred.
- From the example above knowing that a occurs 3 times and that the total
length of the substring is 4 we can tell that there is just 1 different
character (in this case this is b).
- The formula is:
	number of different characters = length of substring - number of highest occurring character
- This means we can always check if we satisfy the k condition without having to
check the number of different characters. all we need is the length of substring
which is equal to (r-l+1, sliding window) and number of highest occurring character.
- Now we are left with the problem of finding the number of highest occurring
character. This is an issue because to find the value we will have to search
the map each time we iterate thereby increasing the time complexity to o(n*2).
- This problem is a maximum problem, hence we don't need to search for each
highest occurring character. We only need the max and to update it when we
get a higher value. This works because we know that if the value isn't
greater than the current max then our result will not change.
- So we create a maxF variable to deal with the issue of finding
number of highest occurring character
- with each loop we will check to update the value of the longest substring
i.e res = max(res, r-l+1). res, is short for result.


Code
- create a function that returns the max integer between two integers

- Create the variables
count - is a map that has byte (character of the string) as key and
number of occurrence as value.
maxF - An integer that stores the max value of occurrence gotten so far
l - the left window (which we will slide)
res - The max length of valid substring

- run a for loop over the length of the string
- add the current character to the map
- Try to update the value of maxF
- Check if we meet the k condition
- if we don't reduce the count of the character
at index l and increment l.
- update the value of res if we have a longer substring

- return res

*NOTE: if you struggle to understand the solution, you can replace the
* use of maxF by actually searching for the value of the highest occurring
* character in the list. You'd have to search for this is in the for loop.
* - also you don't have to get it on the first try. Take a break and
* come back later.

*/
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
