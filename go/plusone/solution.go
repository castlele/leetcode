package plusone

func plusOne(digits []int) []int {
	d := digits

	if len(d) == 0 {
		return []int{1}
	}

	i := len(d) - 1

	var addOne func(int)

	addOne = func(from int) {
		if from < 0 {
			copy := make([]int, len(digits)+1)
			copy = append(d[:1], d[0:]...)
			copy[0] = 1
			d = copy
			return
		}

		if d[from] == 9 {
			d[from] = 0
			addOne(from - 1)
		} else {
			d[from]++
		}
	}

	addOne(i)

	return d
}
