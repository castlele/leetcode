import Testing
@testable import leetcode

@Test
func testRomanToInteger_1() {
    let result = romanToInt("III")
    #expect(result == 3)
}

@Test
func testRomanToInteger_2() {
    let result = romanToInt("IV")
    #expect(result == 4)
}

@Test
func testRomanToInteger_3() {
    let result = romanToInt("IX")
    #expect(result == 9)
}

@Test
func testRomanToInteger_4() {
    let result = romanToInt("LVIII")
    #expect(result == 58)
}

@Test
func testRomanToInteger_5() {
    let result = romanToInt("MCMXCIV")
    #expect(result == 1994)
}
