import Testing

@testable import leetcode

private struct Input {
    let nums: [Int]
    let k: Int
    let expected: [Int]
}

@Test(
    arguments: [
        Input(
            nums: [1, 1, 1, 2, 2, 3],
            k: 2,
            expected: [1, 2]
        ),
        Input(
            nums: [1],
            k: 1,
            expected: [1]
        ),
        Input(
            nums: [1, 2, 1, 2, 1, 2, 3, 1, 3, 2],
            k: 2,
            expected: [1, 2]
        ),
    ]
)
private func topKFrequentTest(input: Input) {
    let result = topKFrequent(input.nums, input.k)

    #expect(counts(result) == counts(input.expected))
}
