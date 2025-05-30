/// source: https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
func strStr(_ haystack: String, _ needle: String) -> Int {
    guard !haystack.isEmpty, !needle.isEmpty, needle.count <= haystack.count else {
        return -1
    }

    outer: for i in 0..<haystack.count - needle.count + 1 {
        var startIndex = haystack.index(haystack.startIndex, offsetBy: i)
        var needleIndex = needle.startIndex
        let endIndex = haystack.index(haystack.startIndex, offsetBy: i + needle.count)

        while startIndex != endIndex {
            if haystack[startIndex] == needle[needleIndex] {
                startIndex = haystack.index(after: startIndex)
                needleIndex = needle.index(after: needleIndex)
            } else {
                continue outer
            }
        }

        return i
    }

    return -1
}
