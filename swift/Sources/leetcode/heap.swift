struct Heap<Element> {
    private var storage: [Element]
    private var comparator: (Element, Element) -> Bool

    init(_ initial: [Element], comparator: @escaping (Element, Element) -> Bool) {
        self.storage = initial
        self.comparator = comparator

        if !storage.isEmpty {
            // n = count - 1
            // n_parent = (n - 1) / 2 = (count - 1 - 1) / 2 = (count - 2) / 2 = count / 2 - 1
            var index = storage.count / 2 - 1

            while index >= 0 {
                shiftDown(from: index)
                index -= 1
            }
        }
    }

    func peek() -> Element? {
        storage.first
    }

    mutating func pop() -> Element? {
        guard let element = peek() else {
            return nil
        }

        storage.swapAt(0, storage.count - 1)
        storage.removeLast()

        if !storage.isEmpty {
            shiftDown(from: 0)
        }

        return element
    }

    mutating func push(_ element: Element) {
        storage.append(element)
        shiftUp(from: storage.count - 1)
    }

    private mutating func shiftUp(from index: Int) {
        var child = index

        while index > 0 {
            let parent = (child - 1) / 2

            guard comparator(storage[child], storage[parent]) else {
                return
            }

            storage.swapAt(child, parent)
            child = parent
        }
    }

    private mutating func shiftDown(from index: Int) {
        var parent = index

        while true {
            let left = 2 * parent + 1
            let right = 2 * parent + 2
            var candidate = parent

            if left < storage.count, comparator(storage[left], storage[candidate]) {
                candidate = left
            }

            if right < storage.count, comparator(storage[right], storage[candidate]) {
                candidate = right
            }

            guard candidate != parent else {
                break
            }

            storage.swapAt(parent, candidate)
            parent = candidate
        }
    }
}
