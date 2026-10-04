/// [Link](https://leetcode.com/problems/product-of-array-except-self/description/)
func productExceptSelf(_ nums: [Int]) -> [Int] {
    var prefix = [Int](repeating: 1, count: nums.count)
    var suffix = [Int](repeating: 1, count: nums.count)
    var result = [Int](repeating: 0, count: nums.count)

    for i in 1..<nums.count {
        prefix[i] = prefix[i - 1] * nums[i - 1]
    }

    for i in stride(from: nums.count - 2, through: 0, by: -1) {
        suffix[i] = suffix[i + 1] * nums[i + 1]
    }

    for i in 0..<nums.count {
        result[i] = prefix[i] * suffix[i]
    }

    return result
}
