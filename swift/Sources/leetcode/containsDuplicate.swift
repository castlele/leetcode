/// (Link)[https://leetcode.com/problems/contains-duplicate/]
/// 1st strategy O(nlogn)+O(n):
/// 1. sort a given array
/// 2. iterate over sorted array
/// 3. check with two indecies if any given value is unique
///
/// 2nd strategy O(n):
/// 1. create a hash map
/// 2. iterate over a given array
/// 3. check if value already in a hash map
func containsDuplicate(_ nums: [Int]) -> Bool {
    guard nums.count >= 2 else {
        return false
    }

    var history = Set<Int>()

    for num in nums {
        if history.contains(num) {
            return true
        }

        history.insert(num)
    }

    return false
}
