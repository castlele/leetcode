/// [Link](https://leetcode.com/problems/move-zeroes/)
func moveZeroes(_ nums: inout [Int]) {
    guard nums.count >= 2 else {
        return
    }

    var lhs = 0

    func swap(from index: Int) -> Int? {
        for rhs in (index + 1)..<nums.count {
            guard nums[rhs] != 0 else {
                continue
            }

            return rhs
        }

        return nil
    }

    while lhs < nums.count {
        if nums[lhs] != 0 {
            lhs += 1
            continue
        }

        guard let rhs = swap(from: lhs) else {
            return
        }

        nums[lhs] = nums[rhs]
        nums[rhs] = 0
    }
}
