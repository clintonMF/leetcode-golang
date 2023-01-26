/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

package main

import "fmt"

/* IDEA behind the solution
-- Ground case
since each type of bracket has opening and closing i.e they come in pairs
for it to be true the length of the string must be an even number.
Hence we will check if the length of string is even and return false if it is not

-- main idea
we will create two variables
1. brackets - which is a map that stores that has it key as the opening bracket and
value as the closing bracket
2. stack - a slice that stores closing bracket

we will loop over the given string.
- for each iteration we will confirm if the character is an opening bracket, if it
is we will add the value of its closing bracket to the stack
else (i.e the character is not an opening bracket, it is a closing bracket) we will
check if the last element in the stack is equal to the closing bracket, if they are
equal we will remove the element from the stack. if they are not equal we will return
false (this is because a wrong bracket was used to close another bracket).
*/
func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	brackets := map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}

	stack := []string{}

	for i := 0; i < len(s); i++ {
		if _, ok := brackets[string(s[i])]; ok {
			// the if statement gets all opening brackets or left brackets
			stack = append(stack, brackets[string(s[i])])
		} else {
			// This deals with all closing brackets or right bracket
			// check for equality of the last element and the right
			if len(stack) == 0 {
				return false
			} else {
				if string(s[i]) != stack[len(stack)-1] {
					return false
				} else {
					// since they are equal remove the value from the stack
					stack = stack[:len(stack)-1]
				}
			}
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("({})"))
}
