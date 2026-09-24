import Testing

@testable import leetcode

struct Input {
    let nums: [Int]
    let target: Int
    let expected: [Int]
}

let commonInputs = [
    Input(nums: [2, 7, 11, 15], target: 9, expected: [0, 1])
]

@Test(arguments: commonInputs)
func squaredTimeComplexity(input: Input) {
    let result = twoSumSquared(input.nums, input.target)

    #expect(result == input.expected)
}

@Test(arguments: commonInputs)
func linearTimeComplexity(input: Input) {
    let result = twoSumLinear(input.nums, input.target)

    #expect(result.count == 2)
    #expect(input.expected.contains(result[0]))
    #expect(input.expected.contains(result[1]))
}
