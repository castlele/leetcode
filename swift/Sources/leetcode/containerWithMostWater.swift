/// [Link](https://leetcode.com/problems/container-with-most-water/)
func maxArea(_ height: [Int]) -> Int {
    var maxArea = 0
    var left = 0
    var right = height.count - 1

    while left < right {
        let leftHeight = height[left]
        let rightHeight = height[right]
        let a = min(leftHeight, rightHeight)
        let b = right - left
        let area = a * b

        maxArea = max(area, maxArea)

        if leftHeight < rightHeight {
            left += 1
        } else {
            right -= 1
        }
    }

    return maxArea
}
