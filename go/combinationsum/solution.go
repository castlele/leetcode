package combinationsum

func combinationSum(candidates []int, target int) [][]int {
    var res [][]int
	var backtrace func([]int, int, int)

	backtrace = func(path []int, start, remaining int) {
		if remaining == 0 {
			copy := append([]int{}, path...)
			res = append(res, copy)
			return
		}

		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			newRemaining := remaining - candidates[i]
			path = append(path, candidates[i])
			backtrace(path, i, newRemaining)
			path = path[:len(path)-1]
		}
	}

	backtrace([]int{}, 0, target)

	return res
}
