package mycalendarone

type MyCalendar struct {
	storage []pair
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

// source: https://leetcode.com/problems/my-calendar-i
func (this *MyCalendar) Book(startTime int, endTime int) bool {
	newTime := p(startTime, endTime)

	if len(this.storage) == 0 {
		this.storage = append(this.storage, newTime)
		return true
	}

	lhs := 0
	rhs := len(this.storage) - 1

	for lhs <= rhs {
		mid := lhs + (rhs-lhs)/2
		cur := this.storage[mid]

		if (newTime.lhs > cur.lhs && newTime.lhs < cur.rhs) ||
			(newTime.rhs > cur.lhs && newTime.rhs < cur.rhs) ||
			newTime.lhs <= cur.lhs && newTime.rhs >= cur.rhs {
			return false
		} else if cur.lhs < newTime.lhs {
			lhs = mid + 1
		} else {
			rhs = mid - 1
		}
	}

	this.storage = append(this.storage, newTime)
	copy(this.storage[lhs+1:], this.storage[lhs:])
	this.storage[lhs] = newTime

	return true
}

type pair struct {
	lhs int
	rhs int
}

func p(lhs, rhs int) pair {
	return pair{lhs, rhs}
}

func (this pair) union(other pair) bool {
	if other.lhs >= this.rhs || this.lhs >= other.rhs {
		return false
	}

	return true
}
