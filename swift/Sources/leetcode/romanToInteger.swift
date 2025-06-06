/// source: https://leetcode.com/problems/roman-to-integer/
func romanToInt(_ s: String) -> Int {
    var result = 0
    var prevValue = Int.min
    var index = s.optIndex(before: s.endIndex)

    while let i = index, i >= s.startIndex {
        let romeNum = s[i]
        let romeValue = alphabet[romeNum] ?? 0

        if prevValue <= romeValue {
            result += romeValue
        } else {
            result -= romeValue
        }

        prevValue = romeValue
        index = s.optIndex(before: i)
    }

    return result
}

extension String {
    func optIndex(before index: String.Index) -> String.Index? {
        index == startIndex ? nil : self.index(before: index)
    }
}

private let alphabet: [Character: Int] = [
    "I": 1,
    "V": 5,
    "X": 10,
    "L": 50,
    "C": 100,
    "D": 500,
    "M": 1000
]
