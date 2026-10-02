/// [Link](https://leetcode.com/problems/top-k-frequent-elements/description)
/// historyTable
/// heap

/// 1. interate over nums collecting freaquency of the num and adding it to the heap
/// 2. get slice of size k from the heap
func topKFrequent(_ nums: [Int], _ k: Int) -> [Int] {
    let historyTable = nums.reduce(into: [:]) { into, item in
        into[item, default: 0] += 1
    }
    var heap = CustomHeap(Array(historyTable.keys)) { lhs, rhs in
        historyTable[lhs, default: 0] > historyTable[rhs, default: 0]
    }

    return (0..<k).compactMap { _ in heap.pop() }
}

private struct CustomHeap {
    private var storage = [Int]()
    private let ordered: (Int, Int) -> Bool

    init(_ array: [Int], sort: @escaping (Int, Int) -> Bool) {
        storage = array
        self.ordered = sort

        if !storage.isEmpty {
            for i in stride(from: (storage.count / 2) - 1, through: 0, by: -1) {
                sink(i)
            }
        }
    }

    func slice(_ k: Int) -> [Int] {
        Array(storage.prefix(k))
    }

    mutating func pop() -> Int? {
        guard !storage.isEmpty else { return nil }
        storage.swapAt(0, storage.count - 1)
        let top = storage.removeLast()
        if !storage.isEmpty { sink(0) }
        return top
    }

    private mutating func sink(_ index: Int) {
        var index = index
        let count = storage.count
        while true {
            let leftChild = 2 * index + 1
            let rightChild = leftChild + 1
            var candidate = index

            if leftChild < count && ordered(storage[leftChild], storage[candidate]) {
                candidate = leftChild
            }
            if rightChild < count && ordered(storage[rightChild], storage[candidate]) {
                candidate = rightChild
            }
            if candidate == index { break }

            storage.swapAt(index, candidate)
            index = candidate
        }
    }
}
