/// [Link](https://leetcode.com/problems/3sum/)
func threeSum(_ nums: [Int]) -> [[Int]] {
    let sorted = nums.sorted()
    var result = [[Int]]()

    func findTwoSum(
        target: Int,
        _ left: inout Int,
        _ right: inout Int
    ) -> (lhsIndex: Int, rhsIndex: Int)? {
        while left < right {
            let sum = sorted[left] + sorted[right]

            if sum == target {
                return (left, right)
            } else if sum > target {
                right -= 1
            } else {
                left += 1
            }
        }

        return nil
    }

    for index in 0..<sorted.count {
        guard index == 0 || sorted[index] != sorted[index - 1] else {
            continue
        }

        var left = index + 1
        var right = sorted.count - 1

        while let (lhsIndex, rhsIndex) = findTwoSum(
            target: -sorted[index],
            &left,
            &right
        ) {
            result.append([sorted[index], sorted[lhsIndex], sorted[rhsIndex]])

            while left < right && sorted[right] == sorted[right - 1] {
                right -= 1
            }

            right -= 1
            left += 1
        }
    }

    return result
}
