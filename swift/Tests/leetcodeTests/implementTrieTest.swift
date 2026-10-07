import Testing

@testable import leetcode

func searchTest() {
    let sut = Trie()

    sut.insert("Hello")

    #expect(sut.search("Hello"))
}

func startsWithTest() {
    let sut = Trie()

    sut.insert("Hello")

    #expect(sut.startsWith("He"))
}
