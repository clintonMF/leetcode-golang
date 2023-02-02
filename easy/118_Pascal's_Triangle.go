/*
Given an integer numRows, return the first numRows of
Pascal's triangle.

In Pascal's triangle, each number is the sum of the two
numbers directly above it as shown:


Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example 2:

Input: numRows = 1
Output: [[1]]


Constraints:

1 <= numRows <= 30
*/

package main

import "fmt"

/* EXPLANATION
IDEA
*NOTE: please go to Leetcode website to view the diagram. It will help
* understand the explanation below.
- From the illustration (on Leetcode) we can see that the element is
equal to the sum of the two elements directly above it in the previous
array. I.e [1 1] => [1 2 1], 2 = 1 + 1
- There is only one exception to this observation and that is the first
and last number in each array which are usually 1. For those case we
assume that there is an imaginary 0 before and after the array.
- Our solution to the problem is to make those imaginary zeros real
we will add the imaginary 0 to the start of the array and to the end
of the array i.e [0] + [1,1] + [0] => [0,1,1,0]
- given that we know [0,1,1,0] => [1,2,1]. from that we can model
the values of the index of the next row to be the expression below
next_row[i] = prev_row[i] + prev_row[i+1]
for example; according to the example above we have
	1 = 0 + 1
	2 = 1 + 1
	1 = 1 + 0

CODE
- We initialize a variable res which is an slice containing a slice of
integers i.e res = [[1]]
- we iterate numsRows times using a for loop
- In the for loop
	- We get the prev slice and call it prev_array
	- we create a new slice called temp which will be the previous array
	and the imaginary zeros
	- we add the imaginary zeros
	- we create a new slice called row (this will hold the element of the
	current row).
	- we loop through temp except the last value (this is to prevent out
		of range error)
	- In the loop
		- according to our model we append temp[j] + temp[j+1] into
		position j in the row

	- The row is completed and appended to res

	- we return res
*/

func generate(numRows int) [][]int {
	res := [][]int{[]int{1}}

	for i := 0; i < numRows-1; i++ {
		prev_array := res[len(res)-1]
		temp := []int{0}
		prev_array = append(prev_array, 0)
		temp = append(temp, prev_array...)
		row := []int{}

		for j := 0; j < len(temp)-1; j++ {
			row = append(row, temp[j]+temp[j+1])
		}

		res = append(res, row)

	}

	return res
}

func main() {

	fmt.Println(generate(5))
}
