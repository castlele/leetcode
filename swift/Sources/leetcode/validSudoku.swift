/// [Link](https://leetcode.com/problems/valid-sudoku/)
func isValidSudoku(_ board: [[Character]]) -> Bool {
    for rowIndex in 0..<9 {
        let nums = board[rowIndex].compactMap {
            $0.isNumber ? Int(String($0)) : nil
        }
        let uniqueNums = Set(nums)

        if nums.count != uniqueNums.count {
            return false
        }
    }

    for columnIndex in 0..<9 {
        var nums = [Int]()

        for rowIndex in 0..<9 {
            let char = board[rowIndex][columnIndex]

            if char.isNumber, let num = Int(String(char)) {
                nums.append(num)
            }
        }

        if nums.count != Set(nums).count {
            return false
        }
    }

    for boxRow in 0..<3 {
        for boxCol in 0..<3 {
            var nums = [Int]()

            for r in 0..<3 {
                for c in 0..<3 {
                    let char = board[boxRow * 3 + r][boxCol * 3 + c]

                    if char.isNumber, let num = Int(String(char)) {
                        nums.append(num)
                    }
                }
            }

            if nums.count != Set(nums).count {
                return false
            }
        }
    }

    return true
}
