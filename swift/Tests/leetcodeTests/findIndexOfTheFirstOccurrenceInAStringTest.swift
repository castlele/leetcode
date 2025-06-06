import Testing
@testable import leetcode

struct FindIndexOfTheFirstOccurrenceInAStringTest {
    @Test
    func test_1() {
        let haystack = "sadbutsad"
        let needle = "sad"

        let result = strStr(haystack, needle)

        #expect(result == 0)
    }

    @Test
    func test_2() {
        let haystack = "sadbutsad"
        let needle = "tsad"

        let result = strStr(haystack, needle)

        #expect(result == 5)
    }

    @Test
    func test_3() {
        let haystack = "leetcode"
        let needle = "leeto"

        let result = strStr(haystack, needle)

        #expect(result == -1)
    }
}
