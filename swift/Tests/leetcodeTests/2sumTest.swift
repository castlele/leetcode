import Testing

@testable import leetcode

private struct Input {
    let nums: [Int]
    let target: Int
    let expected: [Int]
}

private let commonInputs = [
    Input(nums: [2, 7, 11, 15], target: 9, expected: [0, 1])
]

@Test(arguments: commonInputs)
private func squaredTimeComplexity(input: Input) {
    let result = twoSumSquared(input.nums, input.target)

    #expect(result == input.expected)
}

@Test(arguments: commonInputs)
private func linearTimeComplexity(input: Input) {
    let result = twoSumLinear(input.nums, input.target)

    #expect(result.count == 2)
    #expect(input.expected.contains(result[0]))
    #expect(input.expected.contains(result[1]))
}
