import Testing
@testable import leetcode

struct GenerateParenthesesTest {
    @Test
    func test_1() {
        let n = 1

        let result = generateParentheses(n)

        #expect(result == ["()"])
    }

    @Test
    func test_2() {
        let n = 3

        let result = generateParentheses(n)

        #expect(result == ["((()))","(()())","(())()","()(())","()()()"])
    }
}
