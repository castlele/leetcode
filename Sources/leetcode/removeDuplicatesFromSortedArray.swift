/// source: https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(_ nums: inout [Int]) -> Int {
    guard nums.count > 1 else {
        return nums.count
    }

    var index = 1

    for i in 1..<nums.count {
        guard nums[i - 1] != nums[i] else {
            continue
        }

        nums[index] = nums[i]

        index += 1
    }

    nums.removeSubrange(index..<nums.count)

    return nums.count
}
