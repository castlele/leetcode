import Testing

@testable import leetcode

private struct Input {
    let input: [Int]
    let expected: [Int]
}

private let args = [
    Input(
        input: [1, 2, 3, 4],
        expected: [24, 12, 8, 6]
    ),
    Input(
        input: [-1, 1, 0, -3, 3],
        expected: [0, 0, 9, 0, 0]
    ),
]

@Test(arguments: args)
private func testProductExceptSelf(input: Input) {
    let result = productExceptSelf(input.input)

    #expect(result == input.expected)
}
