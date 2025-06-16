package searchinsertpos

func searchInsert(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	if target > nums[r] {
		return r + 1
	}

	if target < nums[l] {
		return l
	}

	var mid int

	for l <= r {
		mid = l + (r-l)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}

type Searcher struct{}

func (s *Searcher) SearchInsert(nums []int, target int) int {
	return searchInsert(nums, target)
}
