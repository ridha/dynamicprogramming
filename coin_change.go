// You are given coins of different denominations and a total amount of money amount.
// Write a function to compute the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.

package dynamicprogramming

import "math"

const INT_MAX int = math.MaxInt32

func CoinChange(coins []int, s int) int {

	// result array stores the minimum number of coins required for s = i
	result := make([]int, s+1)

	for i := range result {
		result[i] = INT_MAX
	}

	// For s = 0
	result[0] = 0

	// Compute values bottom-up
	for i := 1; i <= s; i++ {

		// Go through all coins <= i
		for _, coin := range coins {
			if coin <= i {
				if result[i-coin]+1 < result[i] {
					result[i] = result[i-coin] + 1
				}
			}
		}

	}

	if result[s] == INT_MAX {
		return -1
	}
	return result[s]
}
