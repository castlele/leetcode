import Testing

@testable import leetcode

private struct Input {
    let data: [String]
    let expected: [[String]]
}

@Test(
    arguments: [
        Input(
            data: ["eat", "tea", "tan", "ate", "nat", "bat"],
            expected: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
        ),
        Input(
            data: [""],
            expected: [[""]]
        ),
        Input(
            data: ["", ""],
            expected: [["", ""]]
        ),
        Input(
            data: ["a"],
            expected: [["a"]]
        ),
    ]
)
private func groupAnagramsTest(input: Input) {
    let result = groupAnagrams(input.data)

    #expect(counts(result) == counts(input.expected))
}
