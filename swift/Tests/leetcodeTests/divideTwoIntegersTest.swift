import Testing
@testable import leetcode

struct DivideTwoIntegersTest {
    @Test
    func test_1() {
        let dividend = 10
        let divisor = 3

        let result = divide(dividend, divisor)

        #expect(result == 3)
    }

    @Test
    func test_2() {
        let dividend = 7
        let divisor = -3

        let result = divide(dividend, divisor)

        #expect(result == -2)
    }

    @Test
    func test_3() {
        let dividend = -1
        let divisor = 1

        let result = divide(dividend, divisor)

        #expect(result == -1)
    }

    @Test
    func test_4() {
        let dividend = -1
        let divisor = -1

        let result = divide(dividend, divisor)

        #expect(result == 1)
    }

    @Test
    func test_5() {
        let dividend = -2147483648
        let divisor = 1

        let result = divide(dividend, divisor)

        #expect(result == -2147483648)
    }
}
