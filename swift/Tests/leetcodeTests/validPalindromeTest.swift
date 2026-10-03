import Testing

@testable import leetcode

private struct Input {
    let input: String
    let expected: Bool
}

private let inputs = [
    Input(
        input: "A man, a plan, a canal: Panama",
        expected: true
    ),
    Input(
        input: "ace a car",
        expected: false
    ),
    Input(
        input: " ",
        expected: true
    ),
]

@Test(arguments: inputs)
private func isPalindromeTest(input: Input) {
    let result = isPalindrome(input.input)

    #expect(result == input.expected)
}
