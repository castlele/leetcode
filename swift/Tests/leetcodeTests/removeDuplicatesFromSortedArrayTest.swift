import Testing
@testable import leetcode

struct RemoveDuplicatesFromSortedArrayTest {

    @Test
    func test_1() {
        var nums = [1,1,2]
        let expected = [1, 2]

        let result = removeDuplicates(&nums)

        #expect(result == expected.count)
        #expect(nums == expected)
    }

    @Test
    func test_2() {
        var nums = [0,0,1,1,1,2,2,3,3,4]
        let expected = [0,1,2,3,4]

        let result = removeDuplicates(&nums)

        #expect(result == expected.count)
        #expect(nums == expected)
    }
}
