/// [Link](https://leetcode.com/problems/merge-sorted-array/)
func merge(_ nums1: inout [Int], _ m: Int, _ nums2: [Int], _ n: Int) {
    var last = nums1.count - 1
    var i = m - 1
    var j = n - 1

    while i >= 0 && j >= 0 {
        let lhs = nums1[i]
        let rhs = nums2[j]

        if lhs > rhs {
            nums1[i] = nums1[last]
            nums1[last] = lhs
            i -= 1
        } else {
            nums1[last] = rhs
            j -= 1
        }

        last -= 1
    }

    guard i < 0 else {
        return
    }

    while last >= 0 && j >= 0 {
        nums1[last] = nums2[j]
        last -= 1
        j -= 1
    }
}
