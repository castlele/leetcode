import Testing

@testable import leetcode

private let args = [
    ("leetcode", 0),
    ("loveleetcode", 2),
    ("aabb", -1),
]

@Test(arguments: args)
func firstUniqCharTest(input: (s: String, expected: Int)) {
    let result = firstUniqChar(input.s)

    #expect(result == input.expected)
}
