/// source: https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(_ digits: String) -> [String] {
    guard !digits.isEmpty else { return [] }


    var result: [String] = []

    func backtrack(_ index: Int, _ path: String) {
        if index == digits.count {
            result.append(path)
            return
        }

        let digit = digits[digits.index(digits.startIndex, offsetBy: index)]
        guard let letters = phoneMap[digit] else { return }

        for letter in letters {
            backtrack(index + 1, path + String(letter))
        }
    }

    backtrack(0, "")
    return result
}

private let phoneMap: [Character: [Character]] = [
    "2": ["a", "b", "c"],
    "3": ["d", "e", "f"],
    "4": ["g", "h", "i"],
    "5": ["j", "k", "l"],
    "6": ["m", "n", "o"],
    "7": ["p", "q", "r", "s"],
    "8": ["t", "u", "v"],
    "9": ["w", "x", "y", "z"]
]
