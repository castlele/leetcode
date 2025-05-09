/// Source: https://leetcode.com/problems/palindrome-number/
func isPalindrome(_ x: Int) -> Bool {
    // Negative numbers are not palindromes
    guard x >= 0 else {
        return false
    }

    // Numbers less than 10 are palindromes
    guard x >= 10 else {
        return true
    }

    let str = String(x)
    var s = str.startIndex
    var e = str.index(before: str.endIndex)

    while s < e {
        if str[s] != str[e] {
            return false
        }

        s = str.index(after: s)
        e = str.index(before: e)
    }

    return true
}
