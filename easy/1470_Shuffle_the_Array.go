/*
Given the array nums consisting of 2n elements in the form
[x1,x2,...,xn,y1,y2,...,yn].

Return the array in the form [x1,y1,x2,y2,...,xn,yn].



Example 1:

Input: nums = [2,5,1,3,4,7], n = 3
Output: [2,3,5,4,1,7]
Explanation: Since x1=2, x2=5, x3=1, y1=3, y2=4, y3=7
then the answer is [2,3,5,4,1,7].

Example 2:

Input: nums = [1,2,3,4,4,3,2,1], n = 4
Output: [1,4,2,3,3,2,4,1]


Example 3:

Input: nums = [1,1,2,2], n = 2
Output: [1,2,1,2]


Constraints:

1 <= n <= 500
nums.length == 2n
1 <= nums[i] <= 10^3
*/

package main

/*EXPLANATION

I have provided 2 solutions to this DSA problem.
The first solution is perhaps easier to grasp and more
intuitive. While the second is more efficient

FIRST SOLUTION shuffle1
IDEA
- The idea is to split the array/slice into 2 equal halves
and add their elements to an empty slice according to the
condition given in the question. i.e the first element of
the first array is added, then the first element of the
second array, then the second element of the first array,
then the second element of the second array.

CODE
- Create 2 slices which will be equivalent to the individual
halves of the original slice.
- initialize the value of the original slice to an empty
slice. i.e nums = []int{}.
- Run a for loop over the first slice and get the index i.
	- Append the element of the first slice at index i into nums and
	the element of the second slice at index i into nums
- return nums

SECOND SOLUTION (shuffle)
IDEA
- A close inspection of the shuffle condition reveals that we
are expected to arrange the element according to the index below
1, n+1, 2, n+2, ...., n-1, 2n - 1
- If it is not clear, you can toy with the first solution.
The element at index i in the second array is the element at
index n + i in the first array. You can test this out
CODE
- We create a new array called res which is an array of integers
- we loop through half of the element positions in nums and get
their index (i)
	- We then append the ith element
	- followed by the (n+i)th element
- return res


*/
func shuffle1(nums []int, n int) []int {
	first := nums[:n]
	second := nums[n:]

	nums = []int{}

	for i := 0; i < len(first); i++ {
		nums = append(nums, first[i])
		nums = append(nums, second[i])
	}

	return nums
}

func shuffle(nums []int, n int) []int {
	res := []int{}

	for i := 0; i < len(nums[:n]); i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}

	return res
}
