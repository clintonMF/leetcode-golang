/*
You are given a string title consisting of one or more words separated by a single space,
where each word consists of English letters. Capitalize the string by changing the
capitalization of each word such that:

If the length of the word is 1 or 2 letters, change all letters to lowercase.
Otherwise, change the first letter to uppercase and the remaining letters to lowercase.
Return the capitalized title.



Example 1:

Input: title = "capiTalIze tHe titLe"
Output: "Capitalize The Title"
Explanation:
Since all the words have a length of at least 3, the first letter of each word is uppercase,
and the remaining letters are lowercase.

Example 2:

Input: title = "First leTTeR of EACH Word"
Output: "First Letter of Each Word"
Explanation:
The word "of" has length 2, so it is all lowercase.
The remaining words have a length of at least 3, so the first letter of each remaining
word is uppercase, and the remaining letters are lowercase.

Example 3:

Input: title = "i lOve leetcode"
Output: "i Love Leetcode"
Explanation:
The word "i" has length 1, so it is lowercase.
The remaining words have a length of at least 3, so the first letter of each remaining
 word is uppercase, and the remaining letters are lowercase.


Constraints:

1 <= title.length <= 100
title consists of words separated by a single space without any leading or trailing spaces.
Each word consists of uppercase and lowercase English letters and is non-empty.
*/

package main

import (
	"fmt"
	"strings"
)

/*
EXPLANATION

Ground case
if the length of the string is equal to or less than 2 letters,
return the lowercase of the string. This is in accordance with
the question

Main idea
- first we want to split the string into single words
- In a sentence words are separated by spaces, hence we split based
on this into a variable "arr" which is a slice.
- We iterate over the slice arr, if the length of the word is less
than 2 we change it to lower case, else we change the word to lower case
and make it a title (i.e capitalize the first letter).
- The next step is to create the final string, we want to run the for loop
once, hence we declare a variable "str" which is a string. We add the word
that has been changed in the step above this to the str separated by a
space (i.e " ")

- we finally return str
*/

func capitalizeTitle(title string) string {
	if len(title) <= 2 {
		return strings.ToLower(title)
	}

	arr := strings.Split(title, " ")
	var str string

	for i, v := range arr {
		if len(v) <= 2 {
			v = strings.ToLower(v)
		} else {
			v = strings.ToLower(v)
			v = strings.Title(v)
		}

		if i == 0 {
			str = v
		} else {
			str = str + " " + v
		}
	}

	return str
}

func main() {
	fmt.Println(capitalizeTitle("here To stay"))
}
