/*
Given a non-negative integer x, return the square root of x rounded down to the
nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.


Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to
the nearest integer, 2 is returned.


Constraints:

0 <= x <= 231 - 1
*/

package main

import "fmt"

/* EXPLANATION
* Note : if you are not familiar with the concept of binary search
* please make sure to look the explanation and solution to
* 35_Search_Insert_Position as it applies this concept.

IDEA
- Given that we are not allowed to use built in functions and
exponents. We will have to search for the square root and the
fastest way to do this is using binary search.
- I won't explain binary search here because I have already done
that in 35_Search_Insert_Position. please check it out.
- I will only explain the uniqueness of this question

EXPLANATION
- You may be puzzled on how we are to use binary search when we
were not given an array. You are right, but we know that the square
root of a positive number lies between 0 and the number. This is
where we will carry out our search.
- We will implement binary search with some minor changes
1. To speed up our search we will make mid = low + (low + high)//2
2. We are comparing mid*mid to x
*/

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	if x < 4 {
		return 1
	}

	low := 0
	high := x

	for low <= high {
		mid := low + (high-low)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return high
}

func main() {
	fmt.Println(mySqrt(234))
}
