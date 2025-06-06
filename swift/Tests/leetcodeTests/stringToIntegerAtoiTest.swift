import Testing
@testable import leetcode

@Test
func testStringToInteger_1() {
    let result = myAtoi("42")

    #expect(result == 42)
}

@Test
func testStringToInteger_2() {
    let result = myAtoi("-042")

    #expect(result == -42)
}

@Test
func testStringToInteger_3() {
    let result = myAtoi("1337c0d3")

    #expect(result == 1337)
}

@Test
func testStringToInteger_4() {
    let result = myAtoi("0-1")

    #expect(result == 0)
}

@Test
func testStringToInteger_5() {
    let result = myAtoi("words and 987")

    #expect(result == 0)
}

@Test
func testStringToInteger_6() {
    let result = myAtoi("   -042")

    #expect(result == -42)
}

@Test
func testStringToInteger_7() {
    let result = myAtoi("-91283472332")

    #expect(result == -2147483648)
}
