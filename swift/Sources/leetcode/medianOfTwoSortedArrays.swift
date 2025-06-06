// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(_ nums1: [Int], _ nums2: [Int]) -> Double {
    Solutions().mergeSortSolution(nums1, nums2)
}

private class Solutions {
    func mergeSortSolution(_ nums1: [Int], _ nums2: [Int]) -> Double {
        let m = nums1.count
        let n = nums2.count
        var p1 = 0
        var p2 = 0

        @discardableResult
        func move() -> Int {
            if p1 < m && p2 < n {
                if nums1[p1] < nums2[p2] {
                    p1 += 1
                    return nums1[p1 - 1]

                } else {
                    p2 += 1
                    return nums2[p2 - 1]
                }
            } else if p1 < m {
                p1 += 1
                return nums1[p1 - 1]
            } else {
                p2 += 1
                return nums2[p2 - 1]
            }
        }

        for _ in 0..<((m + n) / 2) - ((m + n) % 2 == 0 ? 1 : 0) {
            move()
        }

        return if (m + n) % 2 == 0 {
            Double(move() + move()) / 2
        } else {
            Double(move())
        }
    }
}
