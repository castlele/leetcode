import Testing

@testable import leetcode

private struct Input {
    let lhs: String
    let rhs: String
    let expected: Bool
}

@Test(
    arguments: [
        Input(
            lhs: "anagram",
            rhs: "nagaram",
            expected: true
        ),
        Input(
            lhs: "rat",
            rhs: "car",
            expected: false
        ),
    ]
)
private func isAnagramTests(input: Input) {
    let result = isAnagram(input.lhs, input.rhs)

    #expect(result == input.expected)
}
