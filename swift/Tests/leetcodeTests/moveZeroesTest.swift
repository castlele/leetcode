import Testing

@testable import leetcode

private let args = [
    ([0, 1, 0, 3, 12], [1, 3, 12, 0, 0]),
    ([0], [0]),
]

@Test(arguments: args)
func moveZeroesTest(input: (nums: [Int], expected: [Int])) {
    var result = input.nums

    moveZeroes(&result)

    #expect(result == input.expected)
}
