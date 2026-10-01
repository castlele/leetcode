/// [Link](https://leetcode.com/problems/group-anagrams/description/)
func groupAnagrams(_ strs: [String]) -> [[String]] {
    var history = [String: [String]]()

    for str in strs {
        let sorted = String(str.sorted())
        history[sorted, default: []].append(str)
    }

    return Array(history.values)
}
