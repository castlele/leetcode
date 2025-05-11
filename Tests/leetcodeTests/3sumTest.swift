import Testing
@testable import leetcode

@Test
func testWithMultipleTriplets() throws {
    let input = [-1, 0, 1, 2, -1, -4]
    let expected: Set<[Int]> = [[-1, -1, 2], [-1, 0, 1]]

    let result = Set(threeSum(input).map { $0.sorted() })

    #expect(result == expected)
}

@Test
func testWithNoTriplets() throws {
    let input = [1, 2, 3, 4, 5]
    let expected: [[Int]] = []

    #expect(threeSum(input) == expected)
}

@Test
func testWithAllZeros() throws {
    let input = [0, 0, 0, 0]
    let expected: [[Int]] = [[0, 0, 0]]

    #expect(threeSum(input) == expected)
}

@Test
func testWithEmptyArray() throws {
    let input: [Int] = []
    let expected: [[Int]] = []

    #expect(threeSum(input) == expected)
}

@Test
func testWithNegativeAndPositive() throws {
    let input = [-2, 0, 1, 1, 2]
    let expected: Set<[Int]> = [[-2, 0, 2], [-2, 1, 1]]

    let result = Set(threeSum(input).map { $0.sorted() })
    #expect(result == expected)
}
