/// source: https://leetcode.com/problems/divide-two-integers/
func divide(_ dividend: Int, _ divisor: Int) -> Int {
    if dividend == Int32.min && divisor == -1 {
        return Int(Int32.max)
    }

    if dividend == Int32.max && divisor == 1 {
        return Int(Int32.max)
    }

    var result = 0
    let isNegative = (divisor < 0 && dividend >= 0) || (divisor >= 0 && dividend < 0)
    let moduleDivisor = abs(divisor)
    var moduleDividend = abs(dividend)

    while moduleDividend >= moduleDivisor {
        var temp = moduleDivisor
        var multiple = 1

        while moduleDividend >= (temp << 1) {
            temp <<= 1
            multiple <<= 1
        }

        moduleDividend -= temp
        result += multiple
    }

    if isNegative {
        result = 0 - result
    }

    return result
}
