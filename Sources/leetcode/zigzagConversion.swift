func convert(_ s: String, _ numRows: Int) -> String {
    var i = 0
    var j = 0
    var isZag = false

    func next() -> (Int, Int) {
        defer {
            if i == numRows - 1 {
                isZag = true
            } else if i == 0 {
                isZag = false
            }

            if isZag {
                i -= 1
                j += 1
            } else {
                i += 1
            }

            if i < 0 {
                i = 0
            }
        }
        return (i, j)
    }

    var result = [[String]](repeating: [String](repeating: "", count: s.count), count: numRows)
    var plainIndex = s.startIndex

    while plainIndex < s.endIndex {
        let char = String(s[plainIndex])
        let (_i, _j) = next()

        result[_i][_j] = char

        plainIndex = s.index(after: plainIndex)
    }

    return result.map { $0.joined() }.joined()
}

