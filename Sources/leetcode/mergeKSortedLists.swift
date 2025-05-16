class ListNode {
    public var val: Int
    public var next: ListNode?
    public init() { self.val = 0; self.next = nil; }
    public init(_ val: Int) { self.val = val; self.next = nil; }
    public init(_ val: Int, _ next: ListNode?) { self.val = val; self.next = next; }
}

/// source: https://leetcode.com/problems/merge-k-sorted-lists/
func mergeKLists(_ lists: [ListNode?]) -> ListNode? {
    guard !lists.isEmpty else { return nil }

    var root: ListNode?

    func merge(other: ListNode?) {
        var new: ListNode?
        var curr: ListNode?
        var lhs = root
        var rhs = other

        while let l = lhs, let r = rhs {
            if l.val < r.val {
                if new == nil {
                    new = l
                    curr = new
                } else {
                    curr?.next = l
                    curr = curr?.next
                }

                lhs = l.next

            } else {
                if new == nil {
                    new = r
                    curr = new
                } else {
                    curr?.next = r
                    curr = curr?.next
                }

                rhs = r.next
            }
        }

        while let n = lhs {
            curr?.next = n
            curr = curr?.next
            lhs = lhs?.next
        }

        while let n = rhs {
            curr?.next = n
            curr = curr?.next
            rhs = rhs?.next
        }

        root = new
    }

    for list in lists {
        if root == nil {
            root = list
        } else {
            guard let list else { continue }

            merge(other: list)
        }
    }

    return root
}

private func extractList(_ head: ListNode?) -> [Int] {
    var result = [Int]()
    var node = head
    while let n = node {
        result.append(n.val)
        node = n.next
    }
    return result
}
