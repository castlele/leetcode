package sqrtx

// source: https://leetcode.com/problems/sqrtx
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	fx := float64(x)
	var lhs float64 = 0
	rhs := fx

	for rhs-lhs > 0 {
		var mid float64 = lhs + (rhs-lhs)/2
		pow := float64(int(mid * mid))

		if pow == fx {
			return int(mid)
		} else if pow < fx {
			lhs = mid
		} else {
			rhs = mid
		}
	}

	return 0
}
