/*
Given two strings s and t, return true if t is an anagram of s, and
false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a
different word or phrase, typically using all the original letters
exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false


Constraints:

1 <= s.length, t.length <= 5 * 104
s and t consist of lowercase English letters.


Follow up: What if the inputs contain Unicode characters?
How would you adapt your solution to such a case?
*/

package main

import "fmt"

/* EXPLANATION
GROUND CASE
- if the length of S and t are different then one cannot make the other (i.e
they don't have the same letters or numbers of letters). so we return false

OTHER CASE
- We want to create two maps, one for s and the other for t. these maps
we have the letters of the string as key, and the number of times they are
repeated as value.
- we will then compare both maps, if they are anagrams they will have the same
keys and value (i.e both maps will be equal)
*/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	dictS := make(map[string]int)
	dictT := make(map[string]int)

	for i, _ := range s {
		if _, ok := dictS[string(s[i])]; ok {
			dictS[string(s[i])] += 1
		} else {
			dictS[string(s[i])] = 1
		}

		if _, ok := dictT[string(t[i])]; ok {
			dictT[string(t[i])] += 1
		} else {
			dictT[string(t[i])] = 1
		}
	}

	for key, _ := range dictS {
		if dictS[key] != dictT[key] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}
