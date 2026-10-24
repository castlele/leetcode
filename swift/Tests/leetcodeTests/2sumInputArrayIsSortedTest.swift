import Testing

@testable import leetcode

private struct Input {
    let nums: [Int]
    let target: Int
    let expected: [Int]
}

private let args = [
    Input(
        nums: [2, 7, 11, 15],
        target: 9,
        expected: [1, 2]
    ),
    Input(
        nums: [2, 3, 4],
        target: 6,
        expected: [1, 3]
    ),
    Input(
        nums: [-1, 0],
        target: -1,
        expected: [1, 2]
    ),
    Input(
        nums: [1, 2, 3, 4, 4, 9, 56, 90],
        target: 8,
        expected: [4, 5]
    ),
]

@Test(arguments: args)
private func twoSumInputArrayIsSortedTest(input: Input) {
    let result = twoSumSorted(input.nums, input.target)

    #expect(result == input.expected)
}
