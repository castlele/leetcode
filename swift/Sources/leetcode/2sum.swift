/*
   Input: nums = [2,7,11,15], target = 9

    O(n^2):
    iterate over the given array with two indices twice
    for (i in 0..<nums.count) {
        a = nums[i]
        for (j in i+1..<nums.count) {
            b = nums[j]
            if (b + a == target) return [i, j]
        }
    }
*/
func twoSumSquared(_ nums: [Int], _ target: Int) -> [Int] {
    for i in 0..<nums.count {
        let leftValue = nums[i]

        for j in (i + 1)..<nums.count {
            let rightValue = nums[j]

            if leftValue + rightValue == target {
                return [i, j]
            }
        }
    }

    fatalError("Invalid input")
}

/*
   Input: nums = [2,7,11,15], target = 9

    O(n):
    map = {value: index}

    for (i in 0..<nums.count) {
        let addend = target - nums[i]
        let secondAddendIndex = map[addend]

        if secondAddendIndex != null {
            return [secondAddend, i]
        } else {
            map[nums[i]] = i
        }
    }
*/
func twoSumLinear(_ nums: [Int], _ target: Int) -> [Int] {
    var map = [Int: Int]()

    for i in 0..<nums.count {
        let firstAddend = nums[i]
        let secondAddend = target - firstAddend

        if let secondAddendIndex = map[secondAddend] {
            return [i, secondAddendIndex]
        } else {
            map[firstAddend] = i
        }
    }

    fatalError("Invalid input")
}
