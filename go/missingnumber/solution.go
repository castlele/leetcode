package missingnumber

// source: https://leetcode.com/problems/missing-number
func missingNumber(nums []int) int {
	n := len(nums)
	sum := 0
	expectedSum := n * (n + 1) / 2

	for _, num := range nums {
		sum += num
	}

	return expectedSum - sum
}
