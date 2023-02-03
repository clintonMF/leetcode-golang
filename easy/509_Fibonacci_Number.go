/*
The Fibonacci numbers, commonly denoted F(n) form a sequence, called
the Fibonacci sequence, such that each number is the sum of the two
preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
Given n, calculate F(n).



Example 1:

Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2:

Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3:

Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.


Constraints:

0 <= n <= 30
*/

package main

import "fmt"

/* the easiest solution to this problem
involves the use of recursion. and it is shown  below

but it's time complexity is terrible O(2**n)

we can make an iterative function that has a better time
complexity, big o of O(n)
*/

/* recursive function
func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

*/

/* iterative function
first we return the base case
f(0) and f(1) are 0 and 1 respectively

for other cases we create 2 variables
n1 - holds the f(o) value 0
n2 - holds the f(1) value 1

for each iteration we increase the value of n
in both n1 and n2 i.e
n1 will now become n2 which is f(1)
and n2 will become n1 + n2 which is f(2)

and so on.
we finally return n2 which is the value of f(n)

*/

func fib(n int) int {
	if n < 2 {
		return n
	}
	n1 := 0
	n2 := 1

	for n > 1 {
		n1, n2 = n2, n1+n2

		n -= 1
	}

	return n2
}

func main() {
	fmt.Println(fib(8))
}
