/// source https://leetcode.com/problems/3sum/
func threeSum(_ nums: [Int]) -> [[Int]] {
    var result = [[Int]]()
    let n = nums.sorted()

    for i in 0..<n.count {
        if i > 0 && n[i] == n[i - 1] { continue }

        let target = -n[i]

        var left = i + 1
        var right = n.count - 1

        while left < right {
            let sum = n[left] + n[right]

            if sum == target {
                result.append([n[i], n[left], n[right]])

                while left < right && n[left] == n[left + 1] {
                    left += 1
                }
                while left < right && n[right] == n[right - 1] {
                    right -= 1
                }

                left += 1
                right -= 1

            } else if sum < target {
                left += 1
            } else {
                right -= 1
            }
        }
    }

    return result
}
