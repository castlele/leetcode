/// source: https://leetcode.com/problems/3sum-closest/description/
func threeSumClosest(_ nums: [Int], _ target: Int) -> Int {
    let nums = nums.sorted()
    var result: Int?

    for i in 0..<nums.count {
        let a = nums[i]

        var left = i + 1
        var right = nums.count - 1

        while left < right {
            let b = nums[left]
            let c = nums[right]
            let sum = a + b + c

            if sum == target {
                return target
            } else if result == nil {
                result = sum
            } else if let r = result, abs(target - sum) < abs(target - r) {
                result = sum
            }

            if sum < target {
                left += 1
            } else {
                right -= 1
            }
        }
    }

    return result ?? 0
}
