import Testing
@testable import leetcode

@Test
func testMedian_1() {
    let result = findMedianSortedArrays([1, 2], [3, 4])

    #expect(result == 2.5)
}

@Test
func testMedian_2() {
    let result = findMedianSortedArrays([1, 3], [2])

    #expect(result == 2)
}
