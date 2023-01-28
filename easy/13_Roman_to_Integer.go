/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12
is written as XII, which is simply X + II. The number 27 is written as XXVII, which
is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However,
the numeral for four is not IIII. Instead, the number four is written as IV. Because
the one is before the five we subtract it making four. The same principle applies to
the number nine, which is written as IX. There are six instances where subtraction
is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.



Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.


Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/

package main

import "fmt"

/* EXPLANATION
IDEA
- We can iterate over the string and lookup the integer value of each
character. The summation of the integer values should give us the
integer of the given string.
- The case above is not always true, there is an exception which is when
a lower integer value comes before (prev) a higher one (curr)
(i.e iv = 4 not 1 + 5 = 6).
we treat such case by adding both and subtracting the lower value twice
(i.e iv = 1 + 5 - 2 * 1 = 4).
- We check for both cases while summing up and this will give us our
integer value.

CODE
- First we store the values of each roman numeral in a dictionary with its
corresponding integer value. We choose a dictionary because it has o(1)
lookup time (instant lookup time) which we make our algorithm faster.
- We create 3 variables
1. Total - keeps track of the integer value of the given string.
2. Prev - The integer value of the previous roman character
3. curr - The integer value of the current roman character
prev and curr enable us to verify which case we are faced with so we can
apply the proper addition.
- We loop over the string, check the character value in our dict and equate
it to curr.
- We confirm what case we are faced with the exception (curr > prev) or the
regular case. If curr > prev we add the curr to the total and subtract the
prev twice (i.e total = total + curr - 2 * prev). Else (i.e curr <= prev),
we increment the total by curr.
- Don't forget to make the prev equal to the curr at the end of the loop.
- return total
*/

func romanToInt(s string) int {
	dict := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var total int
	var prev int
	var curr int

	for _, v := range s {
		curr = dict[string(v)]
		if curr > prev {
			total = total + curr - 2*prev
		} else {
			total = total + curr
		}

		prev = curr
	}

	return total
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
