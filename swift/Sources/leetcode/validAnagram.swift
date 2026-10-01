/// [Link](https://leetcode.com/problems/valid-anagram/description/)
/// 1st O(nlogn) + O(nlogn) + O(n):
/// 1. check strings length
/// 2. iterate over one string and check if other has the same character
///
/// 2nd:
/// 1. create hash table where key is a char and value is a number of occuarances of that char
/// 2. place lhs's characters into this hash table
/// 3. iterate over rhs's characters subtracting values from has table
func isAnagram(_ s: String, _ t: String) -> Bool {
    guard s.count == t.count else {
        return false
    }

    var historyTable = [String.UTF8View.Element: Int]()

    for char in Array(s.utf8) {
        if let count = historyTable[char] {
            historyTable[char] = count + 1
        } else {
            historyTable[char] = 1
        }
    }

    for char in Array(t.utf8) {
        if let count = historyTable[char], count - 1 >= 0 {
            historyTable[char] = count - 1
        } else {
            return false
        }
    }

    return true
}
