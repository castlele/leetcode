package permutations

// source: https://leetcode.com/problems/permutations/description/
func permute(nums []int) [][]int {
	result := make([][]int, 0)

	var backtrace func(current, remaining []int)

	backtrace = func(current, remaining []int) {
		if len(remaining) == 0 {
			result = append(result, append([]int{}, current...))
			return
		}

		for i := range remaining {
			path := append(current, remaining[i])
			newRemaining := append([]int{}, remaining[:i]...)
			newRemaining = append(newRemaining, remaining[i+1:]...)
			backtrace(path, newRemaining)
		}
	}

	backtrace(make([]int, 0), nums)

	return result
}
