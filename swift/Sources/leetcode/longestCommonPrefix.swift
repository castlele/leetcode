/// [Link](https://leetcode.com/problems/longest-common-prefix/description/)
func longestCommonPrefix(_ strs: [String]) -> String {
    let trie = Trie()

    for str in strs {
        guard !str.isEmpty else {
            return ""
        }

        trie.insert(str)
    }

    return trie.longestCommon()
}
