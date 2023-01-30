/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.

You want to maximize your profit by choosing a single day to buy one
stock and choosing a different day in the future to sell that stock.

Return the maximum profit you can achieve from this transaction. If
you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6),
profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because
you must buy before you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.

Constraints:

1 <= prices.length <= 105
0 <= prices[i] <= 104
*/
package main

import (
	"fmt"
	"math"
)

/* EXPLANATION
IDEA
- The trick is knowing the least price so far as we iterate
- If we know the least price up till now, we can figure out the
profit we will make now and check if it is the maxProfit. if it
is we update the value of maxProfit.

CODE
- We create 3 variables
1. buy - A floating point that stores the lowest value at which
we can buy
2. profit - A floating point that stores the profit we can make
at each price
3. max_diff - A floating point that stores the max profit so far
* Note all the variables are all floating points to enable us
* perform mathematical operations

- We iterate over the prices, we check if the current value is
less than buy. if it is we update it (this is done using min func)
- We check if the current profit is greater than max_diff if it
is we update max_diff (this is done using max func)

- We return max_diff
*/

func maxProfit(prices []int) int {
	var buy float64 = float64(prices[0])
	var profit float64 = 0
	var max_diff float64 = 0

	for _, v := range prices {
		v := float64(v)
		buy = math.Min(buy, v)
		profit = v - buy
		max_diff = math.Max(profit, max_diff)
	}

	return int(max_diff)
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
