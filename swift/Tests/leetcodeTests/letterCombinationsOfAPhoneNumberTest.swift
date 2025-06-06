import Testing
@testable import leetcode

struct LetterCombinationsOfAPhoneNumberTest {
    @Test
    func letterCombinations_1() {
        let digits = "23"
        let expected = ["ad","ae","af","bd","be","bf","cd","ce","cf"]

        let result = letterCombinations(digits)

        #expect(isEquals(expected: expected, actual: result))
    }

    @Test
    func letterCombinations_2() {
        let digits = ""
        let expected = [String]()

        let result = letterCombinations(digits)

        #expect(isEquals(expected: expected, actual: result))
    }

    @Test
    func letterCombinations_3() {
        let digits = "2"
        let expected = ["a","b","c"]

        let result = letterCombinations(digits)

        #expect(isEquals(expected: expected, actual: result))
    }
}

extension LetterCombinationsOfAPhoneNumberTest {
    private func isEquals(expected: [String], actual: [String]) -> Bool {
        guard expected.count == actual.count else {
            return false
        }

        for el in expected {
            if !actual.contains(el) {
                return false
            }
        }

        return true
    }
}
