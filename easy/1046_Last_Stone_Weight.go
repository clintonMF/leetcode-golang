/*
You are given an array of integers stones where stones[i]
is the weight of the ith stone.

We are playing a game with the stones. On each turn, we choose
the heaviest two stones and smash them together. Suppose the
heaviest two stones have weights x and y with x <= y. The result
of this smash is:

If x == y, both stones are destroyed, and
If x != y, the stone of weight x is destroyed, and the stone of
weight y has new weight y - x.
At the end of the game, there is at most one stone left.

Return the weight of the last remaining stone. If there are no
stones left, return 0.

Example 1:

Input: stones = [2,7,4,1,8,1]
Output: 1
Explanation:
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of the last stone.
Example 2:

Input: stones = [1]
Output: 1

Constraints:

1 <= stones.length <= 30
1 <= stones[i] <= 1000
*/
package main

import "sort"

/* EXPLANATION
IDEA
- A sorted array will have its highest values as the last
2 elements. If they are equal we can remove them, else we
make the second to the last equal to their difference.
- we can repeat the process above until the length of
our array is less than 2.
- then we return the appropriate answer.

CODE
- run a while loop with the condition that length of stones
greater than 1.
	- sort the values in stone
	- check if the last 2 elements of the sorted array are equal
	- if they are, remove them: this is done by slicing the last
	2 elements out.
	- if they are not, update the value of the last element to
	the difference between the last 2 element. Then slice off
	the last element
- If the length of stones is equal to 1. return the only element
in the array.
- Else return 0.
*/

func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		sort.Ints(stones)
		l := len(stones)
		if stones[l-1] == stones[l-2] {
			stones = stones[:l-2]
		} else if stones[l-1] > stones[l-2] {
			stones[l-2] = stones[l-1] - stones[l-2]
			stones = stones[:l-1]
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}

	return 0
}

func main() {
	print(lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
}
