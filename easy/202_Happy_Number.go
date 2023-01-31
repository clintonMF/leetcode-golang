/*
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

Starting with any positive integer, replace the number by the sum
of the squares of its digits.
Repeat the process until the number equals 1 (where it will stay),
or it loops endlessly in a cycle which does not include 1.
Those numbers for which this process ends in 1 are happy.
Return true if n is a happy number, and false if not.



Example 1:

Input: n = 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1


Example 2:

Input: n = 2
Output: false


Constraints:

1 <= n <= 231 - 1


*/

/* EXPLANATION
There are 2 solutions to this problem in this file
the first saves time but uses large memory
while the other saves memory but takes more time
both rely on a single function. This function
gives us the next number according to the explanation

NEXT NUMBER FUNCTION
IDEA
- The next number function is basically a function that
gives the sum of the squares of tens, unit and hundred
values.
- With this in mind we can create a for loop to extract extract
these values individually. The first loop gives the square
of the unit value, the second that of the tens and so on.
CODE
Create a variable next an integer which holds the value of
the next number.
- run a while loop when the current number (n) is greater than zero
- get the remainder when the current number (n) is divided by 10
- this will give us the unit value
- add the square of this value to the next variable
- update current number (n) to current number - remainder divided by
10.


IDEA (To knowing a happy number or sad number)
- We know a number is happy when one of its next numbers is equal to 1.
- A sad number does not have 1 as one of its next numbers, we can also
tell the number is sad when it repeats one of its next numbers and
starts going in a loop.
- Both solutions follow this idea of a repeated loop for sad number

SOLUTION 1
- This solution is ideal for a Set,  you know once the next number is
repeated that it is a sad number. Since a Set has unique values it is
good for such problem.
- Golang doesn't implement set, so I settled for a map. You can create
a map to store the next numbers as keys.
- If you come across a value thats in your map return false
- If you come across one return True

SOLUTION 2
- makes use of 2 pointers, the first is slow and the other is fast
- Slow gets the next number
- fast gets the next number next number
- since a sad number creates a loop, boot pointers will eventually
have the same value at some point.

*/

package main

import "fmt"

func getNextNum(n int) int {
	var next int
	for n > 0 {
		i := n % 10
		next += i * i
		n = (n - i) / 10
	}

	return next
}

// func isHappy(n int) bool {
// 	dict := make(map[int]int)

// 	_, found := dict[n]

// 	for !found {
// 		dict[n] = 1
// 		n = getNextNum(n)

// 		if n == 1 {
// 			return true
// 		}

// 		_, found = dict[n]
// 	}

// 	return false
// }

func isHappy(n int) bool {
	slow := n
	fast := getNextNum(slow)

	for fast != 1 && slow != fast {
		slow = getNextNum(slow)
		fast = getNextNum(getNextNum(fast))

	}

	return fast == 1
}
func main() {
	fmt.Println(getNextNum(19))
	fmt.Println(isHappy(19))
}
