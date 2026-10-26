import Testing

@testable import leetcode

private let args = [
    (input: [1, 8, 6, 2, 5, 4, 8, 3, 7], output: 49),
    (input: [1, 1], output: 1),
]

@Test(arguments: args)
func maxAreaTest(data: (input: [Int], output: Int)) {
    let result = maxArea(data.input)

    #expect(result == data.output)
}
