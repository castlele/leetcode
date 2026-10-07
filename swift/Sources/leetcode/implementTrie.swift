/// [Link](https://leetcode.com/problems/implement-trie-prefix-tree/description/)
class Trie {
    private let root = Node()

    private class Node {
        private var children = [Character: Node]()

        var isEnd = false

        var count: Int {
            children.count
        }

        var first: (Character, Node)? {
            children.first
        }

        subscript(_ char: Character) -> Node? {
            get {
                children[char]
            }
            set(newValue) {
                children[char] = newValue
            }
        }
    }

    init() {

    }

    func longestCommon() -> String {
        var common = ""

        var node = root

        while node.count == 1, !node.isEnd, let (char, next) = node.first {
            common.append(char)
            node = next
        }

        return common
    }

    func insert(_ word: String) {
        var node = root

        for char in word {
            if let next = node[char] {
                node = next
            } else {
                let next = Node()
                node[char] = next
                node = next
            }
        }

        node.isEnd = true
    }

    func search(_ word: String) -> Bool {
        walk(word)?.isEnd ?? false
    }

    func startsWith(_ prefix: String) -> Bool {
        walk(prefix) != nil
    }

    private func walk(_ word: String) -> Node? {
        var node = root

        for char in word {
            guard let next = node[char] else {
                return nil
            }

            node = next
        }

        return node
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * let obj = Trie()
 * obj.insert(word)
 * let ret_2: Bool = obj.search(word)
 * let ret_3: Bool = obj.startsWith(prefix)
 */
