// Given an integer array with all positive numbers and no duplicates,
// find the number of possible combinations that add up to a positive integer target.

package dynamicprogramming

// Sub-problem: find the number of possible combinations that add up to a number smaller than target
// The overall time complexity is O(len(nums) * target)

func CombinationSum4(nums []int, target int) int {
	result := make([]int, target+1)
	result[0] = 1

	for i := 1; i <= target; i++ {
		sum := 0
		for _, num := range nums {
			if num <= i {
				sum += result[i-num]
			}
		}
		result[i] = sum
	}
	return result[target]
}
