/// [Link](https://leetcode.com/problems/longest-consecutive-sequence/)
func longestConsecutive(_ nums: [Int]) -> Int {
    guard nums.count > 0 else {
        return 0
    }

    guard nums.count >= 2 else {
        return 1
    }

    let sortedNums = nums.sorted()
    var current = 1
    var max = 0

    for i in 1..<sortedNums.count {
        let cur = sortedNums[i]
        let prev = sortedNums[i - 1]
        let diff = abs(cur - prev)
        print("\(prev) - \(cur) - \(diff) - \(current) - \(max)")

        if diff == 1 {
            current += 1
        } else if diff == 0 {
            continue
        } else {
            if current > max {
                max = current
            }

            current = 1
        }
    }

    if current > max {
        max = current
        current = 0
    }

    return max
}
