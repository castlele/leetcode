import Foundation

func myAtoi(_ s: String) -> Int {
    let str = s.trimmingCharacters(in: .whitespaces)
    var sign = 1
    var result = 0
    var i = str.startIndex

    guard !str.isEmpty else {
        return result
    }

    if str[i] == "-" {
        sign = -1
        i = str.index(after: i)
    } else if str[i] == "+" {
        i = str.index(after: i)
    }

    while i < str.endIndex, str[i].isNumber {
        let digit = Int(String(str[i]))!

        if result > (Int(Int32.max) - digit) / 10 {
            return sign == 1 ? Int(Int32.max) : Int(Int32.min)
        }

        result = result * 10 + digit
        i = str.index(after: i)
    }

    return sign * result
}
