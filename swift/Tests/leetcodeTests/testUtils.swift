func counts(_ x: [[String]]) -> [[String]: Int] {
    x.map { $0.sorted() }
        .reduce(into: [:]) { $0[$1, default: 0] += 1 }
}

func counts<T: Hashable & Equatable>(_ x: [T]) -> [T: Int] {
    x.reduce(into: [:]) {
        $0[$1, default: 0] += 1
    }
}
