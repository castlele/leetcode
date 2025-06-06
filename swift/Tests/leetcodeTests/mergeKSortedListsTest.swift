import Testing
@testable import leetcode

struct MergeKSortedListsTest {
    @Test
    func testMergeEmptyLists() {
        let result = mergeKLists([])
        #expect(result == nil)
    }

    @Test
    func testAllEmpty() {
        let result = mergeKLists([nil, nil])
        #expect(result == nil)
    }

    @Test
    func testSingleList() {
        let list = buildList([1, 2, 3])
        let result = mergeKLists([list])
        #expect(result == list)
    }

    @Test
    func testMultipleLists() {
        let l1 = buildList([1, 4, 5])
        let l2 = buildList([1, 3, 4])
        let l3 = buildList([2, 6])
        let result = mergeKLists([l1, l2, l3])
        #expect(extractList(result) == [1, 1, 2, 3, 4, 4, 5, 6])
    }

    @Test
    func testSingleElementLists() {
        let lists = [buildList([5]), buildList([1]), buildList([3])]
        let result = mergeKLists(lists)
        #expect(extractList(result) == [1, 3, 5])
    }
}

extension ListNode: Equatable {
    public static func == (lhs: ListNode, rhs: ListNode) -> Bool {
        var l: ListNode? = lhs
        var r: ListNode? = rhs

        while let lv = l, let rv = r {
            if lv.val != rv.val { return false }
            l = lv.next
            r = rv.next
        }

        return l == nil && r == nil
    }
}

func buildList(_ array: [Int]) -> ListNode? {
    let dummy = ListNode()
    var current = dummy
    for val in array {
        current.next = ListNode(val)
        current = current.next!
    }
    return dummy.next
}

func extractList(_ head: ListNode?) -> [Int] {
    var result = [Int]()
    var node = head
    while let n = node {
        result.append(n.val)
        node = n.next
    }
    return result
}
