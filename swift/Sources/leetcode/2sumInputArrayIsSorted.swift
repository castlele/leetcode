/// [Link](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)
///
/// Goal: O(1) space complexity
///
/// Solution:
/// 1. iterate over numbers
/// 2. subtract target with current number, we have a number to search
/// 3. with binary search search a number in numbers
/// 4. return result if found else move to next number
///
/// Time complexity: O(n*logn)
func twoSumSorted(_ numbers: [Int], _ target: Int) -> [Int] {
    func search(for target: Int, ignoringIndex ignore: Int) -> Int? {
        var left = 0
        var right = numbers.count - 1

        while left <= right {
            let midIndex = (right - left) / 2 + left
            let mid = numbers[midIndex]

            if mid == target && midIndex != ignore {
                return midIndex
            } else if target < mid {
                right = midIndex - 1
            } else {
                left = midIndex + 1
            }
        }

        return nil
    }

    guard numbers.count > 2 else {
        return [1, 2]
    }

    for index in 0..<numbers.count {
        let number = numbers[index]
        let searchTarget = target - number

        if let secondNumberIndex = search(for: searchTarget, ignoringIndex: index) {
            return [index + 1, secondNumberIndex + 1]
        }
    }

    fatalError("Invalid input")
}
