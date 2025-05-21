/// source: https://leetcode.com/problems/remove-element/
func removeElement(_ nums: inout [Int], _ val: Int) -> Int {
    guard !nums.isEmpty else {
        return 0
    }

    nums.sort()

    nums.removeAll { $0 == val }

    return nums.count
}
