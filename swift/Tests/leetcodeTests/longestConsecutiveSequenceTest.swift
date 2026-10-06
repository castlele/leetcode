import Testing

@testable import leetcode

private struct Input {
    let input: [Int]
    let expected: Int
}

private let args = [
    Input(
        input: [100, 4, 200, 1, 3, 2],
        expected: 4
    ),
    Input(
        input: [0, 3, 7, 2, 5, 8, 4, 6, 0, 1],
        expected: 9
    ),
    Input(
        input: [1, 0, 1, 2],
        expected: 3
    ),
]

@Test(arguments: args)
private func longestConsecutiveTest(input: Input) {
    let result = longestConsecutive(input.input)

    #expect(result == input.expected)
}
