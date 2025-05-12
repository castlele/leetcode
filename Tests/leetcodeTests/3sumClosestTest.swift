import Testing
@testable import leetcode

struct ThreeSumClosestTests {
    @Test
    func testPositiveNumbers() throws {
        let nums = [1, 2, 3, 4, 5]
        let target = 10
        #expect(threeSumClosest(nums, target) == 10)
    }

    @Test
    func testNegativeNumbers() throws {
        let nums = [-5, -2, -1, -10]
        let target = -8
        #expect(threeSumClosest(nums, target) == -8)
    }

    @Test
    func testMixedNumbers() throws {
        let nums = [-1, 2, 1, -4]
        let target = 1
        #expect(threeSumClosest(nums, target) == 2)
    }

    @Test
    func testWithDuplicates() throws {
        let nums = [0, 2, 1, -3, 1, 1]
        let target = 1
        #expect(threeSumClosest(nums, target) == 0)
    }

    @Test
    func testLargeTarget() throws {
        let nums = [1, 1, 1, 0]
        let target = 100
        #expect(threeSumClosest(nums, target) == 3)
    }

    @Test
    func testSmallTarget() throws {
        let nums = [1, 1, 1, 0]
        let target = -100
        #expect(threeSumClosest(nums, target) == 2)
    }

    @Test
    func testExactMatchExists() throws {
        let nums = [0, 1, 2]
        let target = 3
        #expect(threeSumClosest(nums, target) == 3)
    }

    @Test
    func testMinimalArraySize() throws {
        let nums = [0, 1, 2]
        let target = 0
        #expect(threeSumClosest(nums, target) == 3)
    }

    @Test
    func testAllZeros() throws {
        let nums = [0, 0, 0, 0]
        let target = 1
        #expect(threeSumClosest(nums, target) == 0)
    }

    @Test
    func testFarFromTarget() throws {
        let nums = [100, 200, 300, 400]
        let target = 50
        #expect(threeSumClosest(nums, target) == 600)
    }

    @Test
    func aLotOfNegatives() {
        let nums = [4,0,5,-5,3,3,0,-4,-5]
        let target = -2

        let result = threeSumClosest(nums, target)

        #expect(result == -2)
    }
}

