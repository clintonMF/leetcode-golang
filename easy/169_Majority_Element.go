/*
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3

Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2


Constraints:

n == nums.length
1 <= n <= 5 * 104
-109 <= nums[i] <= 109
*/

package main

import "fmt"

/* EXPLANATION
* There are 2 solutions to this problem in this file
The first is a more intuitive solution with an okay run time
and space complexity.
The second is a genius solution with a better run time and
space complexity.
*/

/* FIRST SOLUTION EXPLANATION
IDEA: This solution involves making use of Maps. We initialize a
map called dict.
- We loop over the slice of nums and store the numbers in slice
as the key of the map and the number of times the number occurs as
the value. Each time we come across a number we've seen before we
increment. Hence dict has num as key and freq as val
- But this arrangement isn't good, we want the frequency to be the key
and num as the value. This will enable us get the num as soon as we
we know the highest freq occurring number i.e dict1[freq]. we achieve
this by looping over the first dict and switching the storing.
- We also need to know the highest occurring number. We create a
variable highest which will store the freq of the highest occurring
number. In the previous loop ran above, we check if the frequencies
are higher than the variable value. if yes then we make highest = freq.

- Finally we return dict1[highest]
*/

/* SECOND SOLUTION EXPLANATION
IDEA
- This solution is possible because we were assured that the highest
occurring number has a freq > n/2.
- Knowing this, if we increment a variable count each time we saw the
highest occurring number and decrement each time see any other number
the count variable will be above 1.


CODE
- We create 2 variables
1. count - An integer that keeps track of the number of times the
highest occurring num occurs minus the number of times other
num occurs (i.e if majorityElement {count++} else {count--})
2. majorityElement - An integer that stores the majority element

- The majorityElement is initialized to the first element in the
slice. The count variable is initialized to 1; This is because we
note that we have counted the first element in the slice
- We loop over the other element in the slice - nums[1:] (i.e all
others except the first). if the current element is equal to
majorityElement we increment count, if it is not we decrement count
- if count is zero, we change the majority element to the current
element and initialize count to 1.

- Finally we return majorityElement

* NOTE
* If this idea isn't clear yet, try re reading the idea and testing
* manually (without code i.e Yourself) on a slice or array.
*/

// func majorityElement(nums []int) int {
// 	dict := make(map[int]int)
// 	for _, v := range nums {
// 		dict[v] += 1
// 	}

// 	dict1 := make(map[int]int)
// 	highest := 0
// 	for i, v := range dict {
// 		dict1[v] = i
// 		if v > highest {
// 			highest = v
// 		}
// 	}

// 	return dict1[highest]
// }

func majorityElement(nums []int) int {
	count := 1
	majorityElement := nums[0]

	for _, v := range nums[1:] {
		if majorityElement == v {
			count++
		} else {
			count--
		}

		if count == 0 {
			majorityElement = v
			count = 1
		}
	}

	return majorityElement
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 2, 1, 2}))
}
