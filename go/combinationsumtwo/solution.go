package combinationsumtwo

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var backtrace func(int, int)
	sort.Ints(candidates)

	backtrace = func(start int, remaining int) {
		if remaining == 0 {
			res = append(res, append([]int{}, path...))
			return
		}

		if remaining < 0 {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			path = append(path, candidates[i])
			backtrace(i+1, remaining-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrace(0, target)

	return res
}
