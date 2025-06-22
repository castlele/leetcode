package combinations

// source; https://leetcode.com/problems/combinations
func combine(n int, k int) [][]int {
	var res [][]int
	var path []int
	var backtracking func(start int, remaining int)

	backtracking = func(start, remaining int) {
		if remaining == 0 {
			res = append(res, append([]int{}, path...))
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtracking(i+1, remaining-1)
			path = path[:len(path)-1]
		}
	}

	backtracking(1, k)

	return res
}
