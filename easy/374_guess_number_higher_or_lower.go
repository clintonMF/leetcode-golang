/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.


Example 1:

Input: n = 10, pick = 6
Output: 6
Example 2:

Input: n = 1, pick = 1
Output: 1
Example 3:

Input: n = 2, pick = 1
Output: 1


Constraints:

1 <= n <= 231 - 1
1 <= pick <= n
*/

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

package main

func guess(num int) int {
	return 5
}

/* NOTE
* You won't be able to run this code as the guess function defined above
is just a fancy funciton and actually does nothing.
* The real guess() function is defined by leetcode, so you can copy the code
* and test it on the leetcode terminal
*/

/* EXPLANATION
*/

func guessNumber(n int) int {
	low := 1
	high := n
	for low <= high {
		mid := (low + high) / 2
		if guess(mid) == -1 {
			high = mid - 1
		} else if guess(mid) == 1 {
			low = mid + 1
		} else if guess(mid) == 0 {
			return mid
		}
	}

	return -1
}
