/// [Link](https://leetcode.com/problems/first-unique-character-in-a-string/description/)
func firstUniqChar(_ s: String) -> Int {
    guard !s.isEmpty else {
        return -1
    }

    let utfView = Array(s.utf8)
    var history = [String.UTF8View.Element: Int]()

    for char in utfView {
        history[char, default: 0] += 1
    }

    for (index, char) in utfView.enumerated() {
        if history[char] == 1 {
            return index
        }
    }

    return -1
}
