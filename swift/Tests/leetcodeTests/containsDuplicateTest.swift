import Testing

@testable import leetcode

private struct Input {
    let data: [Int]
    let expected: Bool
}

private let args = [
    Input(
        data: [1, 2, 3, 1],
        expected: true
    ),
    Input(
        data: [1, 2, 3, 4],
        expected: false
    ),
    Input(
        data: [1, 1, 1, 3, 3, 4, 3, 2, 4, 2],
        expected: true
    ),
]

@Test(arguments: args)
private func containsDuprlicateTest(input: Input) {
    let result = containsDuplicate(input.data)

    #expect(result == input.expected)
}
