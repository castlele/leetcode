/// [Link](https://leetcode.com/problems/valid-palindrome/description/)
func isPalindrome(_ s: String) -> Bool {
    func isAlphanumeric(_ s: Character) -> Bool {
        s.isLetter || s.isNumber
    }

    let str = Array(s)
    var l = 0
    var r = str.count - 1

    while l < r {
        guard isAlphanumeric(str[l]) else {
            l += 1
            continue
        }

        guard isAlphanumeric(str[r]) else {
            r -= 1
            continue
        }

        if str[l].lowercased() != str[r].lowercased() {
            return false
        }

        l += 1
        r -= 1
    }

    return true
}
