import Testing
@testable import leetcode

@Test
func testPalindromeNumber_1() {
    let result = isPalindrome(121)

    #expect(result == true)
}

@Test
func testPalindromeNumber_2() {
    let result = isPalindrome(-121)

    #expect(result == false)
}

@Test
func testPalindromeNumber_3() {
    let result = isPalindrome(10)

    #expect(result == false)
}

@Test
func testPalindromeNumber_4() {
    let result = isPalindrome(1234321)

    #expect(result == true)
}

@Test
func testPalindromeNumber_5() {
    let result = isPalindrome(0)

    #expect(result == true)
}

@Test
func testPalindromeNumber_6() {
    let result = isPalindrome(1)

    #expect(result == true)
}
