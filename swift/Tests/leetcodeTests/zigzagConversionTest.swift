import Testing
@testable import leetcode

@Test
func testZigzagConversion_1() {
    let result = convert("PAYPALISHIRING", 3)

    #expect(result == "PAHNAPLSIIGYIR")
}

func testZigzagConversion_2() {
    let result = convert("AB", 1)

    #expect(result == "AB")
}
