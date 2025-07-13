package twosum

// source: https://leetcode.com/problems/two-sum/
func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	dict := make(map[int]int)

	for i, num := range nums {
		if val, ok := dict[target-num]; ok {
			return []int{val, i}
		}

		dict[num] = i
	}

	return nil
}
