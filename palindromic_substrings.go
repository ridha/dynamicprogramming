// Given a string, your task is to count how many palindromic substrings in this string.
// The substrings with different start indexes or end indexes are counted as different
// substrings even they consist of same characters.

// The input string length won't exceed 1000.

package dynamicprogramming

// Time complexity: O(n^2).
// Use a matrix to check if s[i,j] is a palindrome. If so, add one to count.

func countSubstrings(s string) int {
	count := 0
	n := len(s)

	result := make([][]bool, n)
	for i := range result {
		result[i] = make([]bool, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			if s[i] == s[j] {
				if j-i < 2 {
					result[i][j] = true
				} else {
					result[i][j] = result[i+1][j-1]
				}
				if result[i][j] == true {
					count++
				}
			}
		}
	}
	return count
}
