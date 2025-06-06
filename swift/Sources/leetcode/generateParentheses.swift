/// source: https://leetcode.com/problems/generate-parentheses/description/
func generateParentheses(_ n: Int) -> [String] {
    var result = [String]()

    func backtrack(_ path: String, _ open: Int, _ close: Int) {
        if path.count == n * 2 {
            result.append(path)
            return
        }

        if open < n {
            backtrack(path + "(", open + 1, close)
        }

        if close < open {
            backtrack(path + ")", open, close + 1)
        }
    }

    backtrack("", 0, 0)

    return result
}
