import Testing
@testable import leetcode

struct RemoveElementTest {
    @Test
    func test_1() {
        var nums = [3,2,2,3]
        let val = 3
        let expected = [2, 2]

        let result = removeElement(&nums, val)

        #expect(result == expected.count)
        #expect(nums.sorted() == expected)
    }
}
