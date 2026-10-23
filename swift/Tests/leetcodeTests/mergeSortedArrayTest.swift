import Testing

@testable import leetcode

private struct Input {
    let num1: [Int]
    let m: Int
    let num2: [Int]
    let n: Int
    let expected: [Int]
}

private let args = [
    Input(
        num1: [1, 2, 3, 0, 0, 0],
        m: 3,
        num2: [2, 5, 6],
        n: 3,
        expected: [1, 2, 2, 3, 5, 6]
    ),
    Input(
        num1: [1],
        m: 1,
        num2: [],
        n: 0,
        expected: [1]
    ),
    Input(
        num1: [0],
        m: 0,
        num2: [1],
        n: 1,
        expected: [1]
    ),
]

@Test(arguments: args)
private func mergeTwoArraysTest(input: Input) {
    var result = input.num1

    merge(&result, input.m, input.num2, input.n)

    #expect(result == input.expected)
}
