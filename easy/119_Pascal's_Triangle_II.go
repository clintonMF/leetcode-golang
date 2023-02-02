/*
Given an integer Index, return the Indexth (0-indexed)  of
the Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly
above it as shown:




Example 1:

Input: Index = 3
Output: [1,3,3,1]


Example 2:

Input: Index = 0
Output: [1]


Example 3:

Input: Index = 1
Output: [1,1]


Constraints:

0 <= Index <= 33
*/

package main

import "fmt"

/* EXPLANATION
- It the same explanation as 118_Pascal's_Triangle.go
- The difference is we return the single row needed
- This is implemented by returning the last element in the
array

*/

func get(Index int) []int {
	res := [][]int{[]int{1}}

	for i := 0; i < Index; i++ {
		temp := []int{0}
		arr := res[len(res)-1]
		arr = append(arr, 0)
		temp = append(temp, arr...)
		row := []int{}
		for j := 0; j < len(temp)-1; j++ {
			row = append(row, temp[j]+temp[j+1])
		}

		res = append(res, row)

	}

	return res[Index]
}

func main() {
	fmt.Println(get(3))
}
