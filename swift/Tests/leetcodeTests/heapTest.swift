import Testing

@testable import leetcode

let tests = [
    [0, 9, 8, 7, 6, 5],
    [234, 2, 1, 31304, 12, 21],
    [],
    [-1, 2, -1234, 1234],
    [0],
    [5, 4, 3, 2, 1],
    [3, 1, 3, 1, 2],
]

@Test(arguments: tests)
private func minHeapTest(input: [Int]) {
    let sut = Heap(input, comparator: <)

    #expect(sut.peek() == input.min())
}

@Test(arguments: tests)
private func maxHeapTest(input: [Int]) {
    let sut = Heap(input, comparator: >)

    #expect(sut.peek() == input.max())
}

@Test(arguments: tests)
private func popHeapTest(input: [Int]) {
    var sut = Heap(input, comparator: <)

    var popped = [Int]()
    while let element = sut.pop() {
        popped.append(element)
    }

    #expect(popped == input.sorted())
    #expect(sut.peek() == nil)
}

@Test(arguments: tests)
private func pushHeapTest(input: [Int]) {
    var sut = Heap([Int](), comparator: <)

    for (index, element) in input.enumerated() {
        sut.push(element)

        #expect(sut.peek() == input[...index].min())
    }
}

@Test(arguments: tests)
private func pushThenPopHeapTest(input: [Int]) {
    var sut = Heap([Int](), comparator: <)

    for element in input {
        sut.push(element)
    }

    var popped = [Int]()
    while let element = sut.pop() {
        popped.append(element)
    }

    #expect(popped == input.sorted())
}
