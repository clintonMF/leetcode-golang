/*
Given an array of strings strs, group the anagrams together. You
can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters
of a different word or phrase, typically using all the original
letters exactly once.



Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:

Input: strs = [""]
Output: [[""]]

Example 3:

Input: strs = ["a"]
Output: [["a"]]


Constraints:

1 <= strs.length <= 104
0 <= strs[i].length <= 100
strs[i] consists of lowercase English letters.
*/

package main

import "fmt"

func getKey(str string) string {
	var key = make([]int, 26)
	for _, ch := range str {
		key[ch-'a']++
	}
	return fmt.Sprint(key)
}

func groupAnagrams(strs []string) [][]string {
	var dict map[string][]string = make(map[string][]string)
	for _, word := range strs {
		key := getKey(word)
		dict[key] = append(dict[key], word)
	}

	var res [][]string
	for _, val := range dict {
		res = append(res, val)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"cat", "tac", "act", "tabs"}))
}
