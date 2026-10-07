import Testing

@testable import leetcode

private struct Input {
    let input: [String]
    let expected: String
}

private let args = [
    Input(
        input: ["flower", "flow", "flight"],
        expected: "fl"
    ),
    Input(
        input: ["dog", "racecar", "car"],
        expected: ""
    ),
    Input(
        input: ["", "b"],
        expected: ""
    ),
    Input(
        input: ["ab", "a"],
        expected: "a"
    ),
]

@Test(arguments: args)
private func longestCommonProfixTest(input: Input) {
    let result = longestCommonPrefix(input.input)

    #expect(result == input.expected)
}
